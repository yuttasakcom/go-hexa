package router

import (
	"github.com/yuttasakcom/go-hexa/app/database"
	"github.com/yuttasakcom/go-hexa/app/todo"
)

func RegisterTodoRouter(r *App, db *database.Store) {
	todoModel := todo.NewTodoModel(db)
	handler := todo.NewTodoHandler(todoModel)
	r.Post("/todos", handler.Create)
}
