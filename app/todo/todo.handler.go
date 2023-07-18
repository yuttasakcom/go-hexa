package todo

import (
	"net/http"
)

type TodoHandler struct {
	model modeler
}

func NewTodoHandler(model modeler) *TodoHandler {
	return &TodoHandler{model: model}
}

type Context interface {
	Bind(v interface{}) error
	Status(code int) Context
	JSON(v interface{}) error
}

func (t *TodoHandler) Create(c Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.Status(http.StatusBadRequest).JSON(TodoError{Msg: "Invalid request payload"})
		return
	}

	err := t.model.Create(&todo)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(TodoError{Msg: err.Error()})
		return
	}

	c.Status(http.StatusCreated).JSON(todo)
}
