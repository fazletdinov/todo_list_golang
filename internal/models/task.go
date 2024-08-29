package models

import (
	"time"
)

type Task struct {
	ID          int       `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"not null; size:256" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	DueDate     time.Time `gorm:"not null" json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
