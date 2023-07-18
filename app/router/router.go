package router

import (
	"github.com/yuttasakcom/go-hexa/app/ctx"
	"github.com/yuttasakcom/go-hexa/app/database"
)

type App struct {
	*FiberApp
}

func NewApp() *App {
	return &App{FiberApp: NewFiberApp()}
}

func Register(app *App, db *database.Store) *App {
	// Readiness Probe
	health := NewFiberHandler(func(c ctx.Context) {
		_ = c.Status(200).SendString("OK")
	})
	app.Get("/health", health)

	RegisterTodoRouter(app, db)
	return app
}
