package repository

import (
	"context"
	"time"
	"todo-list/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type taskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		database: db,
	}
}

func (tr *taskRepository) Create(
	ctx context.Context,
	title string,
	description string,
	dueDate time.Time,
) (*models.Task, error) {
	task := models.Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}
	result := tr.database.WithContext(ctx).Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *taskRepository) GetByIDTask(
	ctx context.Context,
	taskID int,
) (*models.Task, error) {
	var task models.Task
	result := tr.database.WithContext(ctx).First(&task, "id = ?", taskID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *taskRepository) GetTasks(
	ctx context.Context,
	limit uint,
	offset uint,
) (*[]models.Task, error) {
	var tasks []models.Task
	result := tr.database.WithContext(ctx).
		Model(&models.Task{}).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tasks, nil
}

func (tr *taskRepository) Update(
	ctx context.Context,
	taskID int,
	title string,
	description string,
	dueDate time.Time,
) (*models.Task, error) {
	var task models.Task
	result := tr.database.WithContext(ctx).
		Model(&task).
		Clauses(clause.Returning{}).
		Where("ID = ?", taskID).
		Updates(models.Task{Title: title, Description: description, DueDate: dueDate, UpdatedAt: time.Now()})
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *taskRepository) Delete(
	ctx context.Context,
	taskID int,
) error {
	result := tr.database.WithContext(ctx).Delete(&models.Task{}, taskID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
