package service

import (
	"context"
	"time"
	"todo-list/internal/domain/repository"
	"todo-list/internal/dto"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(
	taskRepository repository.TaskRepository,
) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (ts *taskService) CreateTask(
	ctx context.Context,
	title string,
	description string,
	dueDate time.Time,
) (*dto.TaskResponse, error) {
	taskResponse, err := ts.taskRepository.Create(ctx, title, description, dueDate)
	if err != nil {
		return nil, err
	}
	return &dto.TaskResponse{
		ID:          taskResponse.ID,
		Title:       taskResponse.Title,
		Description: taskResponse.Description,
		DueDate:     dto.DueDate(taskResponse.DueDate),
		CreatedAt:   taskResponse.CreatedAt,
		UpdatedAt:   taskResponse.UpdatedAt,
	}, nil
}

func (ts *taskService) GetByIDTask(
	ctx context.Context,
	taskID int,
) (*dto.TaskResponse, error) {
	taskResponse, err := ts.taskRepository.GetByIDTask(ctx, taskID)
	if err != nil {
		return nil, err
	}
	return &dto.TaskResponse{
		ID:          taskResponse.ID,
		Title:       taskResponse.Title,
		Description: taskResponse.Description,
		DueDate:     dto.DueDate(taskResponse.DueDate),
		CreatedAt:   taskResponse.CreatedAt,
		UpdatedAt:   taskResponse.UpdatedAt,
	}, nil
}

func (ts *taskService) GetTasks(
	ctx context.Context,
	limit uint,
	offset uint,
) (*[]dto.TaskResponse, error) {
	tasksResponse, err := ts.taskRepository.GetTasks(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	tasks := make([]dto.TaskResponse, 0, limit)
	for _, task := range *tasksResponse {
		tasks = append(tasks, dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     dto.DueDate(task.DueDate),
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return &tasks, nil
}

func (ts *taskService) UpdateTask(
	ctx context.Context,
	taskID int,
	title string,
	description string,
	dueDate time.Time,
) (*dto.TaskResponse, error) {
	taskResponse, err := ts.taskRepository.Update(ctx, taskID, title, description, dueDate)
	if err != nil {
		return nil, err
	}
	return &dto.TaskResponse{
		ID:          taskResponse.ID,
		Title:       taskResponse.Title,
		Description: taskResponse.Description,
		DueDate:     dto.DueDate(taskResponse.DueDate),
		CreatedAt:   taskResponse.CreatedAt,
		UpdatedAt:   taskResponse.UpdatedAt,
	}, nil

}

func (ts *taskService) DeleteTask(
	ctx context.Context,
	taskID int,
) error {
	return ts.taskRepository.Delete(ctx, taskID)
}
