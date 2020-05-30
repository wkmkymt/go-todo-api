package router

import (
	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/infrastructure/api/handler"
)

// NewRouter is Creating Routing
func NewRouter(g *gin.Engine, handler handler.AppHandler) {
	// API Ver.1
	v1Routes := g.Group("/api/v1")
	{
		// Todos
		todoRoutes := v1Routes.Group("/todos")
		{
			todoRoutes.GET("/", handler.TodoHandler.GetAllTodos)
		}
	}
}
