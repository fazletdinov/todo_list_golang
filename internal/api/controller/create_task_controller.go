package controller

import (
	"net/http"
	"time"
	// "time"

	"github.com/gin-gonic/gin"

	"todo-list/config"
	"todo-list/internal/domain/service"
	"todo-list/internal/dto"
)

type CreateTaskController struct {
	TaskService service.TaskService
	Env         *config.Config
}

// CreateTask	godoc
// @Summary		Создание Task
// @Tags        Tasks
// @Accept		json
// @Produce     json
// @Param		body	    body		dto.TaskRequest		   true	    "Создание Task"
// @Success     201  		{object}  	dto.TaskResponse
// @Failure		400			{object}	dto.ErrorResponse
// @Failure		500			{object}	dto.ErrorResponse
// @Router      /tasks 		[post]
func (ct *CreateTaskController) Create(ctx *gin.Context) {
	var request dto.TaskRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных"})
		return
	}

	// dueDate, err := time.Parse(time.RFC3339, request.DueDate)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных"})
	// 	return
	// }

	task, err := ct.TaskService.CreateTask(
		ctx,
		request.Title,
		request.Description,
		time.Time(request.DueDate),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Проблема на сервере"})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}
