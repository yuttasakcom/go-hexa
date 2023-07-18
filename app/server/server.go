package server

import (
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
	router.Register(app, s.db)
	app.Listen(s.cfg.App().Host())
}
