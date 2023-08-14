package router

import (
	"net/http"

	"github.com/yuttasakcom/go-hexa/src/core/adapter"
	"github.com/yuttasakcom/go-hexa/src/core/app"
	"github.com/yuttasakcom/go-hexa/src/core/common"
	"github.com/yuttasakcom/go-hexa/src/core/database"
	"github.com/yuttasakcom/go-hexa/src/domain/todo"
)

func Register(app *app.App, store *database.Store) {
	health := adapter.NewHandler(func(ch common.Ctx) {
		ch.Status(http.StatusOK).JSON(map[string]string{"status": "ok"})
	})
	app.Get("/system/health", health)

	todo.NewTodoRouter(app, store)
}
