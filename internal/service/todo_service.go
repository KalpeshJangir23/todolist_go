package services

import (
	"fmt"

	"github.com/KalpeshJangir23/todolist_go/internal/models"
	"github.com/KalpeshJangir23/todolist_go/internal/repos"
)

type TodoService interface {
	Create(todo models.Todo) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	GetByID(id uint) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

type todoService struct {
	repo repos.TodoRepository
}

func NewTodoService(r repos.TodoRepository) TodoService {
	return &todoService{repo: r}
}

func (s *todoService) Create(todo models.Todo) (models.Todo, error) {
	// Business logic example: title cannot be empty
	if todo.Title == "" {
		return todo, fmt.Errorf("title cannot be empty")
	}
	return s.repo.Create(todo)
}

func (s *todoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetByID(id uint) (models.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *todoService) Update(todo models.Todo) (models.Todo, error) {
	return s.repo.Update(todo)
}

func (s *todoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
