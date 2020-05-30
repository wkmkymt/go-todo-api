package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/infrastructure/api/router"
	"github.com/wkmkymt/go-todo-api/registry"
)

func main() {
	g := gin.Default()

	r := registry.NewAppInteractor()
	h := r.NewAppHandler()

	router.NewRouter(g, h)

	g.Run(":5000")
}
