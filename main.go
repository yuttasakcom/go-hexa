package main

import (
	"os"

	"github.com/yuttasakcom/go-hexa/app/config"
	"github.com/yuttasakcom/go-hexa/app/router"
	"github.com/yuttasakcom/go-hexa/app/store"
	"github.com/yuttasakcom/go-hexa/app/todo"
)

func main() {
	cfg := config.NewConfig(envPath())
	db := store.GormStoreConnect(cfg.Db())
	r := router.NewFiberRouter()
	gormStore := store.NewGormStore(db)
	handler := todo.NewTodoHandler(gormStore)
	r.Post("/todos", handler.Create)
	r.Listen(cfg.App().Host())
}

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
