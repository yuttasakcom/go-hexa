package todo

import (
	context "context"
	"fmt"
)

type todoService struct {
}

func NewTodoService() CreateTodoServiceServer {
	return todoService{}
}

func (todoService) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	result := fmt.Sprintf("Hello %v", req.Title)
	res := &CreateTodoResponse{Result: result}
	return res, nil
}

func (todoService) mustEmbedUnimplementedCreateTodoServiceServer() {}
