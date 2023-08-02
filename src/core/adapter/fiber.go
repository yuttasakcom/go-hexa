package adapter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
)

type FiberApp struct {
	*fiber.App
}

func NewFiberApp() *FiberApp {
	return &FiberApp{fiber.New(fiber.Config{
		CaseSensitive:         false,
		StrictRouting:         false,
		DisableStartupMessage: true,
		ReadBufferSize:        10240,
		ReadTimeout:           30 * time.Second,
	})}
}

func NewFiberHandler(handler func(common.Ctx)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	}
}

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberContext) Status(code int) common.Ctx {
	c.Ctx.Status(code)
	return c
}

func (c *FiberContext) JSON(v interface{}) error {
	return c.Ctx.JSON(v)
}

func (c *FiberContext) Next() error {
	return c.Ctx.Next()
}

func (c *FiberContext) Request() *fasthttp.Request {
	return c.Ctx.Request()
}

func (c *FiberContext) Context() *fasthttp.RequestCtx {
	return c.Ctx.Context()
}

func (c *FiberContext) Method() string {
	return c.Ctx.Method()
}

func (c *FiberContext) Route() *fiber.Route {
	return c.Ctx.Route()
}

func (c *FiberContext) OriginalURL() string {
	return c.Ctx.OriginalURL()
}

func (c *FiberContext) IP() string {
	return c.Ctx.IP()
}

func (c *FiberContext) Protocol() string {
	return c.Ctx.Protocol()
}

func (c *FiberContext) Path() string {
	return c.Ctx.Path()
}

func (c *FiberContext) Locals(key string, value ...interface{}) (val interface{}) {
	if len(value) > 0 {
		c.Ctx.Locals(key, value[0])
	}
	return c.Ctx.Locals(key)
}

func (c *FiberContext) Response() *fasthttp.Response {
	return c.Ctx.Response()
}

func (r *FiberApp) Post(path string, handler func(common.Ctx)) {
	r.App.Post(path, NewFiberHandler(handler))
}
