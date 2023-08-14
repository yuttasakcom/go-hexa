package todo

import (
	"net/http"

	"github.com/yuttasakcom/go-hexa/src/core/common"
	"github.com/yuttasakcom/go-hexa/src/core/middleware"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("todo_usecase")

type TodoHandler struct {
	repo repoer
}

func NewTodoHandler(repo repoer) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (t *TodoHandler) Create(c common.Ctx) {
	_, span := tracer.Start(middleware.GetSpanContext(c), "todo.usecase.Create")
	defer span.End()

	var todo TodoEntity
	if err := c.Bind(&todo); err != nil {
		c.Status(http.StatusBadRequest).JSON(TodoError{Msg: "Invalid request payload"})
		return
	}

	err := t.repo.Create(&todo)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(TodoError{Msg: err.Error()})
		return
	}

	c.Status(http.StatusCreated).JSON(todo)
}
