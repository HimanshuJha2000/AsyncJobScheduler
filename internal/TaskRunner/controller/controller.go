package controller

import (
	"github.com/DevtronLabs/headoutProj/internal/TaskRunner/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type AsyncJobController struct {
	AsynJobService service.Service
}

var AsyncJob AsyncJobController

func (asj *AsyncJobController) StartJob(ctx *gin.Context) {
	nStr := ctx.Params.ByName("sleep_time")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Failed to get sleep time for this task",
			"internal_error ": "Provided Sleep time is invalid",
		})
		return
	}

	statusCode, result, err := asj.AsynJobService.StartJobImpl(n)

	if err != nil {
		log.Println("Error while creating task ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (asj *AsyncJobController) PauseJob(ctx *gin.Context) {
	taskID := ctx.Params.ByName("task_id")

	statusCode, result, err := asj.AsynJobService.PauseJobImpl(taskID)

	if err != nil {
		log.Println("Error while pausing task ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (asj *AsyncJobController) ResumeJob(ctx *gin.Context) {
	taskID := ctx.Params.ByName("task_id")

	statusCode, result, err := asj.AsynJobService.ResumeJobImpl(taskID)

	if err != nil {
		log.Println("Error while resuming task ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (asj *AsyncJobController) TerminateJob(ctx *gin.Context) {
	taskID := ctx.Params.ByName("task_id")

	statusCode, result, err := asj.AsynJobService.TerminateJobImpl(taskID)

	if err != nil {
		log.Println("Error while terminating task ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (asj *AsyncJobController) StatusJob(ctx *gin.Context) {
	taskID := ctx.Params.ByName("task_id")

	statusCode, result, err := asj.AsynJobService.JobStatusImpl(taskID)

	if err != nil {
		log.Println("Error while fetching status of task ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}
