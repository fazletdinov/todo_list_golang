package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo-list/config"
)

func SetupTaskRouter(
	gin *gin.Engine,
	env *config.Config,
	db *gorm.DB,
) {
	publicRouter := gin.Group("/api/v1")
	NewGetTaskRouter(publicRouter, env, db)
	NewCreateTaskRouter(publicRouter, env, db)
	NewDeleteTaskRouter(publicRouter, env, db)
	NewUpdateTaskRouter(publicRouter, env, db)
}
