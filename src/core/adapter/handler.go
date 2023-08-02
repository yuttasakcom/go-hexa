package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
)

func NewHandler(handler func(common.Ctx)) fiber.Handler {
	return NewFiberHandler(handler)
}
