package todo

import "github.com/yuttasakcom/go-kafka-simple/src/core/database"

type repoer interface {
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
