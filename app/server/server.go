package server

import (
	"github.com/yuttasakcom/go-hexa/app/config"
	"github.com/yuttasakcom/go-hexa/app/database"
	"github.com/yuttasakcom/go-hexa/app/router"
	"github.com/yuttasakcom/go-hexa/app/todo"
)

type IServer interface {
	Start()
}

type server struct {
	cfg config.IConfig
	db  *database.Store
	app *router.App
}

func NewServer(cfg config.IConfig, db *database.Store) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: router.NewApp(),
	}
}

func (s *server) Start() {
	todoModel := todo.NewTodoModel(s.db)
	handler := todo.NewTodoHandler(todoModel)
	s.app.Post("/todos", handler.Create)
	s.app.Listen(s.cfg.App().Host())
}
