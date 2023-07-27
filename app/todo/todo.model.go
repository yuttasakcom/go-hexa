package todo

import (
	"github.com/yuttasakcom/go-hexa/app/database"
)

// var collectionName = "todos"

type modeler interface {
	Create(todo *Todo) error
}

type todoModel struct {
	db *database.Store
}

func NewTodoModel(db *database.Store) *todoModel {
	return &todoModel{db: db}
}

func (t *todoModel) Create(todo *Todo) error {
	return t.db.Create(todo)
}
