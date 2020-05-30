package controllers

import (
	"github.com/wkmkymt/go-todo-api/domain/model"
	"github.com/wkmkymt/go-todo-api/usecase/service"
)

type todoController struct {
	TodoService service.TodoService
}

// TodoController is Todo Controller
type TodoController interface {
	GetAllTodos() (model.Todos, error)
}

// NewTodoController is Creating New Controller
func NewTodoController(us service.TodoService) TodoController {
	return &todoController{TodoService: us}
}

func (uc *todoController) GetAllTodos() (model.Todos, error) {
	todos, err := uc.TodoService.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
