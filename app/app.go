package app

import "github.com/yuttasakcom/go-hexa/app/adapter"

type IApp interface {
	Listen(addr string) error
	Shutdown() error
}

type App struct {
	*adapter.FiberApp
}

func NewApp() *App {
	return &App{adapter.NewFiberApp()}
}

func (a *App) Listen(addr string) error {
	return a.App.Listen(addr)
}

func (a *App) Shutdown() error {
	return a.App.Shutdown()
}
