package server

import (
	"github.com/yuttasakcom/go-hexa/app/config"
	"github.com/yuttasakcom/go-hexa/app/database"
	"github.com/yuttasakcom/go-hexa/app/router"
	"github.com/yuttasakcom/go-hexa/app/todo"
	"gorm.io/gorm"
)

type IServer interface {
	Start()
}

type server struct {
	cfg config.IConfig
	db  *gorm.DB
	app *router.FiberRouter
}

func NewServer(cfg config.IConfig, db *gorm.DB) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: router.NewFiberRouter(),
	}
}

func (s *server) Start() {
	gormStore := database.NewGormStore(s.db)
	handler := todo.NewTodoHandler(gormStore)
	s.app.Post("/todos", handler.Create)
	s.app.Listen(s.cfg.App().Host())
}
