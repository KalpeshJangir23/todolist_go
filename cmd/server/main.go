package main

import (
	"log"

	"github.com/KalpeshJangir23/todolist_go/internal/config/db"
	"github.com/KalpeshJangir23/todolist_go/internal/handler"
	"github.com/KalpeshJangir23/todolist_go/internal/repos"
	services "github.com/KalpeshJangir23/todolist_go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	database := db.Connect()

	// Init layers
	repo := repos.NewTodoRepository(database)
	service := services.NewTodoService(repo)
	handler := handler.NewTodoHandler(service)

	// Gin Router
	r := gin.Default()

	// Todo routes
	r.POST("/todos", handler.Create)
	r.GET("/todos", handler.GetAll)
	r.GET("/todos/:id", handler.GetByID)
	r.PUT("/todos/:id", handler.Update)
	r.DELETE("/todos/:id", handler.Delete)

	log.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8080")
}
