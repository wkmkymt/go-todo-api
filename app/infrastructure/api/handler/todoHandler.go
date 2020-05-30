package handler

import (
	"github.com/wkmkymt/go-todo-api/interface/controllers"
)

type todoHandler struct {
	TodoController controllers.TodoController
}

// TodoHandler is Todo Handler
type TodoHandler interface {}

// NewTodoHandler is Creating New Todo Handler
func NewTodoHandler(uc controllers.TodoController) TodoHandler {
	return &todoHandler{TodoController: uc}
}
