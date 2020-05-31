package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/interface/controllers"
)

type todoHandler struct {
	TodoController controllers.TodoController
}

// TodoHandler is Todo Handler
type TodoHandler interface {
	GetAllTodos(c *gin.Context)
}

// NewTodoHandler is Creating New Todo Handler
func NewTodoHandler(uc controllers.TodoController) TodoHandler {
	return &todoHandler{TodoController: uc}
}

func (uh *todoHandler) GetAllTodos(c *gin.Context) {
	todos, err := uh.TodoController.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todos)
}
