package todo

import (
	"context"

	store "github.com/yuttasakcom/go-hexa/app/database"
)

type ITodoUseCase interface {
	Create(c context.Context, todo *Todo) error
	Update(c context.Context, todo *Todo) error
	Delete(c context.Context, todo *Todo) error
	GetByID(c context.Context, todoID int) (Todo, error)
	Get(c context.Context) ([]Todo, error)
}

type TodoUseCase struct {
	store store.Store
}

func NewTodoUseCase(store store.Store) *TodoUseCase {
	return &TodoUseCase{store: store}
}

func (uc *TodoUseCase) Create(c context.Context, todo *Todo) error {
	return uc.store.Create(c, todo)
}

func (uc *TodoUseCase) Update(c context.Context, todo *Todo) error {
	return uc.store.Update(c, todo)
}

func (uc *TodoUseCase) Delete(c context.Context, todo *Todo) error {
	return uc.store.Delete(c, todo)
}

func (uc *TodoUseCase) GetByID(c context.Context, todoID int) (Todo, error) {
	todo, err := uc.store.GetByID(c, Todo{}, todoID)
	if err != nil {
		return Todo{}, err
	}
	return todo.(Todo), nil
}

func (uc *TodoUseCase) Get(c context.Context) ([]Todo, error) {
	todos, err := uc.store.Get(c, Todo{})
	if err != nil {
		return []Todo{}, err
	}
	return uc.convertTodos(todos), nil
}

func (uc *TodoUseCase) convertTodos(todos []interface{}) []Todo {
	var result []Todo
	for _, todo := range todos {
		result = append(result, todo.(Todo))
	}
	return result
}
