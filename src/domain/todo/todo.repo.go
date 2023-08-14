package todo

import "github.com/yuttasakcom/go-hexa/src/core/database"

type repoer interface {
	Create(todo *TodoEntity) error
}

type todoModel struct {
	db *database.Store
}

func NewTodoModel(db *database.Store) *todoModel {
	return &todoModel{db: db}
}

func (t *todoModel) Create(todo *TodoEntity) error {
	return t.db.Create(todo)
}
