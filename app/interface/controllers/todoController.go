package controllers

import (
	"github.com/wkmkymt/go-todo-api/usecase/service"
)

type todoController struct {
	TodoService service.TodoService
}

// TodoController is Todo Controller
type TodoController interface {
}
