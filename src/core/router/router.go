package router

import (
	"net/http"

	"github.com/yuttasakcom/go-kafka-simple/src/core/adapter"
	"github.com/yuttasakcom/go-kafka-simple/src/core/app"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
	"github.com/yuttasakcom/go-kafka-simple/src/core/database"
	"github.com/yuttasakcom/go-kafka-simple/src/domain/todo"
)

func Register(app *app.App, store *database.Store) {
	health := adapter.NewHandler(func(ch common.Ctx) {
		ch.Status(http.StatusOK).JSON(map[string]string{"status": "ok"})
	})
	app.Get("/system/health", health)

	todo.NewTodoRouter(app, store)
}
