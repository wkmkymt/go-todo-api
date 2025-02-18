package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/domain/model"
	"github.com/wkmkymt/go-todo-api/interface/controllers"
)

type todoHandler struct {
	TodoController controllers.TodoController
}

// TodoHandler is Todo Handler
type TodoHandler interface {
	CreateTodo(c *gin.Context)
	GetAllTodos(c *gin.Context)
}

// NewTodoHandler is Creating New Todo Handler
func NewTodoHandler(uc controllers.TodoController) TodoHandler {
	return &todoHandler{TodoController: uc}
}

func (uh *todoHandler) CreateTodo(c *gin.Context) {
	todo := &model.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := uh.TodoController.CreateTodo(todo)
	if err == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create Todo!"})
	}
	c.JSON(http.StatusCreated, todo)
}

func (uh *todoHandler) GetAllTodos(c *gin.Context) {
	todos, err := uh.TodoController.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todos)
}
