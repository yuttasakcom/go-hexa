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

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	app := fiber.New()
	return &FiberRouter{App: app}
}

func (r *FiberRouter) Post(path string, handler func(todo.Context)) {
	r.App.Post(path, NewFiberHandler(handler))
}
