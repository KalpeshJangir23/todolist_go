package repos

import (
	"github.com/KalpeshJangir23/todolist_go/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	GetByID(id uint) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *todoRepository) GetByID(id uint) (models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return todo, err
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) Update(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error
	return todo, err
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
