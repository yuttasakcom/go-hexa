package router

import "github.com/yuttasakcom/go-hexa/app/database"

type App struct {
	*FiberApp
}

func NewApp() *App {
	return &App{FiberApp: NewFiberApp()}
}

func Register(app *App, db *database.Store) *App {
	RegisterTodoRouter(app, db)
	return app
}
