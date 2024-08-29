package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo-list/config"
	"todo-list/internal/api/controller"
	"todo-list/internal/domain/repository"
	"todo-list/internal/domain/service"
)

func NewGetTaskRouter(
	group *gin.RouterGroup,
	env *config.Config,
	db *gorm.DB,
) {
	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := &controller.GetTaskController{
		TaskService: taskService,
		Env:         env,
	}
	group.GET("/tasks/:task_id", taskController.Fetch)
	group.GET("/tasks", taskController.Fetchs)
}
