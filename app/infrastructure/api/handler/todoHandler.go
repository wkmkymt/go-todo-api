package handler

import (
	"github.com/wkmkymt/go-todo-api/interface/controllers"
)

type todoHandler struct {
	TodoController controllers.TodoController
}

// TodoHandler is Todo Handler
type TodoHandler interface {
}
