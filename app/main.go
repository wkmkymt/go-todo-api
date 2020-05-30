package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wkmkymt/go-todo-api/config"
	"github.com/wkmkymt/go-todo-api/infrastructure/api/router"
	"github.com/wkmkymt/go-todo-api/registry"
)

func main() {
	// Load Config
	config.LoadConfig()

	// Get Gin App Engine
	g := gin.Default()

	// Get App Handler
	r := registry.NewAppInteractor()
	h := r.NewAppHandler()

	// Set App Router
	router.NewRouter(g, h)

	// Run Server
	g.Run(":5000")
}
