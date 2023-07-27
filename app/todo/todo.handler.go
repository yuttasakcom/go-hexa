package todo

import (
	"net/http"

	"github.com/yuttasakcom/go-hexa/app/common"
)

type TodoHandler struct {
	model modeler
}

func NewTodoHandler(model modeler) *TodoHandler {
	return &TodoHandler{model: model}
}

func (t *TodoHandler) Create(c common.Context) {
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
