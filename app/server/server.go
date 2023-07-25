package server

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yuttasakcom/go-hexa/app/config"
	"github.com/yuttasakcom/go-hexa/app/database"
	"github.com/yuttasakcom/go-hexa/app/router"
)

type IServer interface {
	Start()
}

type Server struct {
	cfg config.IConfig
	db  *database.Store
}

func NewServer(cfg config.IConfig, db *database.Store) IServer {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) Start() {
	app := router.NewApp()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer wg.Done()
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Shutdown()
	}()

	go func() {
		defer wg.Done()
		router.Register(app, s.db)
		app.Listen(s.cfg.App().Host())
	}()

	wg.Wait()
}
