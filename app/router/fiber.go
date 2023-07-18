package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-hexa/app/todo"
)

type MyContext struct {
	*fiber.Ctx
}

func NewMyContext(c *fiber.Ctx) *MyContext {
	return &MyContext{Ctx: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *MyContext) Status(code int) todo.Context {
	c.Ctx.Status(code)
	return c
}

func (c *MyContext) JSON(v interface{}) error {
	return c.Ctx.JSON(v)
}

func NewFiberHandler(handler func(todo.Context)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler(NewMyContext(c))
		return nil
	}
}

type FiberApp struct {
	*fiber.App
}

func NewFiberApp() *FiberApp {
	app := fiber.New()
	return &FiberApp{App: app}
}

func (r *FiberApp) Post(path string, handler func(todo.Context)) {
	r.App.Post(path, NewFiberHandler(handler))
}

type FiberRouter struct {
	fiber.Router
}

func NewFiberRouter() *FiberRouter {
	return &FiberRouter{}
}
