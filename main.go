package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todo-list/docs"

	"todo-list/internal/api/router"
	"todo-list/internal/app"
	"todo-list/pkg/logger"
)

// @title           Gin Tasks Service
// @version         1.0
// @description     Сервис для создания задач

// @contact.name   Идель Фазлетдинов
// @contact.email  fvi-it@mail.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
	app := app.App()
	gin := gin.Default()
	gin.Use(logger.Logger(app.Log))

	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.SetupTaskRouter(gin, app.Env, app.DB)

	gin.Run(":" + app.Env.TasksServer.TasksPort)
}
