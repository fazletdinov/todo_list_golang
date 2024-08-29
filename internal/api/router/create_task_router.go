package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo-list/config"
	"todo-list/internal/api/controller"
	"todo-list/internal/domain/repository"
	"todo-list/internal/domain/service"
)

func NewCreateTaskRouter(
	group *gin.RouterGroup,
	env *config.Config,
	db *gorm.DB,
) {
	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := &controller.CreateTaskController{
		TaskService: taskService,
		Env:         env,
	}
	group.POST("/tasks", taskController.Create)
}
