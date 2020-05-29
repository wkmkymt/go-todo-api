package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/infrastructure/api/router"
)

func main() {
	g := gin.Default()

	router.NewRouter(g)

	g.Run(":5000")
}
