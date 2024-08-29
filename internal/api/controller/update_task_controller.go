package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"todo-list/config"
	"todo-list/internal/domain/service"
	"todo-list/internal/dto"
)

type UpdateTaskController struct {
	TaskService service.TaskService
	Env         *config.Config
}

// UpdateTask   godoc
// @Summary     Обновление Task
// @Tags        Tasks
// @Accept		json
// @Produce     json
// @Param	    task_id			path		string			     true    "Task ID"
// @Param		body		    body		dto.TaskRequest  	 true	 "Для обновления Task"
// @Success     200  		    {object}  	dto.SuccessResponse
// @Failure	  	400			    {object}	dto.ErrorResponse
// @Failure	  	500			    {object}	dto.ErrorResponse
// @Router      /tasks/{task_id}/ 	[put]
func (utc *UpdateTaskController) Update(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	taskIDParse, err := strconv.Atoi(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных"})
		return
	}

	if _, err = utc.TaskService.GetByIDTask(
		ctx,
		taskIDParse,
	); err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Задача не найдена"})
		return
	}

	var taskRequest dto.TaskRequest

	if err = ctx.ShouldBindJSON(&taskRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных"})
		return
	}
	task, err := utc.TaskService.UpdateTask(
		ctx,
		taskIDParse,
		taskRequest.Title,
		taskRequest.Description,
		time.Time(taskRequest.DueDate),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Проблема на сервере"})
		return
	}

	ctx.JSON(http.StatusOK, task)

}
