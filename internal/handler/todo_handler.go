package handler

import (
	"net/http"
	"strconv"

	"github.com/KalpeshJangir23/todolist_go/internal/models"
	services "github.com/KalpeshJangir23/todolist_go/internal/service"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service services.TodoService
}

func NewTodoHandler(s services.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

// Create Todo
func (h *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.Create(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get Todo by ID
func (h *TodoHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	todo, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// Get All Todos
func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// Update Todo
func (h *TodoHandler) Update(c *gin.Context) {
	id := c.Param("id") // keep it as string (UUID)

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assign UUID string to todo.ID
	todo.ID = id

	updated, err := h.service.Update(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Todo
func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
