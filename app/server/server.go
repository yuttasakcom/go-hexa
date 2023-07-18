package server

import (
	"fmt"
	"os"
	"os/signal"

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
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	router.Register(app, s.db)
	app.Listen(s.cfg.App().Host())
}
