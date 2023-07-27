package todo

import (
	"github.com/yuttasakcom/go-hexa/app"
	"github.com/yuttasakcom/go-hexa/app/database"
)

func NewTodoRouter(r *app.App, db *database.Store) {
	todoModel := NewTodoModel(db)
	handler := NewTodoHandler(todoModel)
	r.Post("/todos", handler.Create)
}
