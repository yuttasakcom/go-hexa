package todo

import (
	"github.com/yuttasakcom/go-hexa/src/core/app"
	"github.com/yuttasakcom/go-hexa/src/core/database"
)

func NewTodoRouter(r *app.App, store *database.Store) {
	todoModel := NewTodoModel(store)
	handler := NewTodoHandler(todoModel)
	r.Post("/todos", handler.Create)
}
