package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yuttasakcom/go-kafka-simple/src/core/adapter"
	"github.com/yuttasakcom/go-kafka-simple/src/core/app"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
	"github.com/yuttasakcom/go-kafka-simple/src/core/database"
	"github.com/yuttasakcom/go-kafka-simple/src/core/middleware"
	"github.com/yuttasakcom/go-kafka-simple/src/core/router"
)

type Serverer interface {
	Start()
}

type Server struct {
	config config.Configer
	store  *database.Store
}

func NewServer(config config.Configer) *Server {
	return &Server{
		config: config,
		store:  database.NewStore(config.DBConfig()),
	}
}

func (s *Server) Start() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

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
	}()

	go func() {
		defer wg.Done()
		router.Register(app, s.store)
		host := s.config.App().Host()
		app.Listen(host)
	}()

	wg.Wait()
}
