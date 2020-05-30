package service

import (
	"github.com/wkmkymt/go-todo-api/usecase/repository"
)

type todoService struct {
	TodoRepository repository.TodoRepository
}

// TodoService is Todo Service
type TodoService interface {}

// NewTodoService is Creating New Todo Service
func NewTodoService(ur repository.TodoRepository) TodoService {
	return &todoService{TodoRepository: ur}
}
