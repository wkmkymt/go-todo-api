package service

import (
	"github.com/wkmkymt/go-todo-api/domain/model"
	"github.com/wkmkymt/go-todo-api/usecase/repository"
)

type todoService struct {
	TodoRepository repository.TodoRepository
}

// TodoService is Todo Service
type TodoService interface {
	Create(todo *model.Todo) bool
	GetAll() (model.Todos, error)
}

// NewTodoService is Creating New Todo Service
func NewTodoService(ur repository.TodoRepository) TodoService {
	return &todoService{TodoRepository: ur}
}

func (us *todoService) Create(todo *model.Todo) bool {
	return us.TodoRepository.Store(todo)
}

func (us *todoService) GetAll() (model.Todos, error) {
	todos, err := us.TodoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
