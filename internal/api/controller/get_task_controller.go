package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"todo-list/config"
	"todo-list/internal/domain/service"
	"todo-list/internal/dto"
)

type GetTaskController struct {
	TaskService service.TaskService
	Env         *config.Config
}

// GetTask	   godoc
// @Summary	   Получение Task
// @Tags       Tasks
// @Accept	   json
// @Produce    json
// @Param	   task_id		      path		    string		          	true		"Task ID"
// @Success    200  		      {object}  	dto.TaskResponse
// @Failure	   400	              {object}	    dto.ErrorResponse
// @Failure	   500			      {object}	    dto.ErrorResponse
// @Router     /tasks/{task_id}    [get]
func (tc *GetTaskController) Fetch(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	taskIDParse, err := strconv.Atoi(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных."})
		return
	}

	task, err := tc.TaskService.GetByIDTask(
		ctx,
		taskIDParse,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Task не найден"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// ListTasks godoc
//
// @Summary		Получение списока Task
// @Tags	    Tasks
// @Accept	    json
// @Produce		json
// @Param	    limit			query				int		true	"limit"
// @Param	    offset			query				int		true	"offset"
// @Success		200	{array}		dto.TaskResponse
// @Failure		400	{object}	dto.ErrorResponse
// @Failure		500	{object}	dto.ErrorResponse
// @Router	    /tasks 			[get]
func (tc *GetTaskController) Fetchs(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных."})
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных."})
		return
	}

	tasks, err := tc.TaskService.GetTasks(ctx, uint(limit), uint(offset))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Task не найден"})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
