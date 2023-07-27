package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-hexa/app/common"
)

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberContext) Status(code int) common.Context {
	c.Ctx.Status(code)
	return c
}

func (c *FiberContext) JSON(v interface{}) error {
	return c.Ctx.JSON(v)
}

func NewFiberHandler(handler func(common.Context)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
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

func (r *FiberApp) Post(path string, handler func(common.Context)) {
	r.App.Post(path, NewFiberHandler(handler))
}
