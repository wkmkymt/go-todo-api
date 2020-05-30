package router

import (
	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/infrastructure/api/handler"
)

// NewRouter is Setting Routing
func NewRouter(g *gin.Engine, handler handler.AppHandler) {}
