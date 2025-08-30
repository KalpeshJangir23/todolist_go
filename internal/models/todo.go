package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Status      string         `gorm:"type:varchar(20);not null;default:'todo'" json:"status"`
	Priority    string         `gorm:"type:varchar(20);not null;default:'medium'" json:"priority"`
	DueDate     *time.Time     `json:"dueDate"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	CompletedAt *time.Time     `json:"completedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
