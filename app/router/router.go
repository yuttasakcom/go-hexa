package router

type App struct {
	*FiberApp
}

func NewApp() *App {
	return &App{FiberApp: NewFiberApp()}
}
