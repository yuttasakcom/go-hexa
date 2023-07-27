package router

import (
	"github.com/yuttasakcom/go-hexa/app"
	"github.com/yuttasakcom/go-hexa/app/adapter"
	"github.com/yuttasakcom/go-hexa/app/common"
	"github.com/yuttasakcom/go-hexa/app/database"
	"github.com/yuttasakcom/go-hexa/app/todo"
)

func Register(app *app.App, db *database.Store) {
	// Readiness Probe
	health := adapter.NewFiberHandler(func(c common.Context) {
		_ = c.Status(200).SendString("OK")
	})
	app.Get("/health", health)

	todo.NewTodoRouter(app, db)
}
