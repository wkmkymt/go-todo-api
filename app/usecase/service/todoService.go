package service

import (
	"github.com/wkmkymt/go-todo-api/usecase/repository"
)

type todoService struct {
	TodoRepository repository.TodoRepository
}

// TodoService is Todo Service
type TodoService interface {
}
