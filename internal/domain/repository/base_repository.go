package repository

import (
	"context"
	"time"
	"todo-list/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, title string, description string, dueDate time.Time) (*models.Task, error)
	GetByIDTask(ctx context.Context, taskID int) (*models.Task, error)
	GetTasks(ctx context.Context, limit uint, offset uint) (*[]models.Task, error)
	Update(ctx context.Context, taskID int, title string, description string, dueDate time.Time) (*models.Task, error)
	Delete(ctx context.Context, taskID int) error
}
