package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"todo-list/config"
	"todo-list/internal/domain/service"
	"todo-list/internal/dto"
)

type DeleteTaskController struct {
	TaskService service.TaskService
	Env         *config.Config
}

// DeleteTask	godoc
// @Summary		Удаление Task
// @Tags        Tasks
// @Accept		json
// @Produce     json
// @Param	    task_id			path		string		          		true 	"Task ID"
// @Success     204  		    {object}  	dto.SuccessResponse
// @Failure		400			    {object}	dto.ErrorResponse
// @Failure		500			    {object}	dto.ErrorResponse
// @Router      /tasks/{task_id} [delete]
func (dtc *DeleteTaskController) Delete(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	taskIDParse, err := strconv.Atoi(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Неправильный формат данных."})
		return
	}

	if _, err = dtc.TaskService.GetByIDTask(
		ctx,
		taskIDParse,
	); err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Задача не найдена"})
		return
	}

	if err = dtc.TaskService.DeleteTask(
		ctx,
		taskIDParse,
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Проблема на сервере"})
		return
	}

	ctx.JSON(http.StatusNoContent, dto.SuccessResponse{Message: "Задача удалена."})

}
