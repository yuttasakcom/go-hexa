package server

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yuttasakcom/go-hexa/src/core/adapter"
	"github.com/yuttasakcom/go-hexa/src/core/app"
	"github.com/yuttasakcom/go-hexa/src/core/common"
	"github.com/yuttasakcom/go-hexa/src/core/config"
	"github.com/yuttasakcom/go-hexa/src/core/database"
	"github.com/yuttasakcom/go-hexa/src/core/middleware"
	"github.com/yuttasakcom/go-hexa/src/core/router"
	"github.com/yuttasakcom/go-hexa/src/domain/todo"
	"google.golang.org/grpc"
)

type Serverer interface {
	Start()
}

type server struct {
	config config.Configer
	store  *database.Store
}

func NewServer(config config.Configer) Serverer {
	return &server{
		config: config,
		store:  database.NewStore(config.DBConfig()),
	}
}

func (s *server) Start() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	grpcServer := grpc.NewServer()

	app := app.NewApp()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowHeaders:  "*",
		AllowMethods:  "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		ExposeHeaders: "content-disposition",
	}))

	app.Use(adapter.NewHandler(func(c common.Ctx) {
		middleware.Tracer(c)
		c.Next()
	}))

	go func() {
		defer wg.Done()
		<-c
		slog.L().Info("Gracefully shutting down...")
		app.Shutdown()
		grpcServer.Stop()
	}()

	go func() {
		defer wg.Done()
		router.Register(app, s.store)
		host := s.config.App().HOST()
		slog.L().Info("Server is running at " + host)
		app.Listen(host)

	}()

	go func() {
		defer wg.Done()
		grpcHost := s.config.App().GRPC_HOST()
		listener, _ := net.Listen("tcp", grpcHost)
		todo.RegisterCreateTodoServiceServer(grpcServer, todo.NewTodoService())
		slog.L().Info("GRPC is running at " + grpcHost)
		err := grpcServer.Serve(listener)
		if err != nil {
			slog.L().Error(err.Error())
		}
	}()

	wg.Wait()
}
