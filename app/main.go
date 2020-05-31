package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/wkmkymt/go-todo-api/config"
	"github.com/wkmkymt/go-todo-api/infrastructure/api/router"
	"github.com/wkmkymt/go-todo-api/infrastructure/datastore"
	"github.com/wkmkymt/go-todo-api/registry"
)

func main() {
	// Load Config
	config.LoadConfig()

	// Get Gin App Engine
	g := gin.Default()

	// Get DB
	db := datastore.NewDB()

	// Get App Handler
	r := registry.NewAppInteractor(db)
	h := r.NewAppHandler()

	// Set App Router
	router.NewRouter(g, h)

	// Stop DB
	defer db.Close()

	// Run Server
	g.Run(":5000")
}
