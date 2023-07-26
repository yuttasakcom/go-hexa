package todo

import (
	"context"
	"time"

	"github.com/yuttasakcom/go-hexa/app/database"
)

var collectionName = "todos"

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := t.db.Collection(collectionName).InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}
