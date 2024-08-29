package service

import (
	"context"
	"time"
	"todo-list/internal/dto"
)

type TaskService interface {
	CreateTask(ctx context.Context, title string, description string, dueDate time.Time) (*dto.TaskResponse, error)
	GetByIDTask(ctx context.Context, taskID int) (*dto.TaskResponse, error)
	GetTasks(ctx context.Context, limit uint, offset uint) (*[]dto.TaskResponse, error)
	UpdateTask(ctx context.Context, taskID int, title string, description string, dueDate time.Time) (*dto.TaskResponse, error)
	DeleteTask(ctx context.Context, taskID int) error
}
