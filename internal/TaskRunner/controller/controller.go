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
			"error":           "Failed to get Image by this ImageID",
			"internal_error ": "Cat Image ID is not provided",
		})
		return
	}

	statusCode, result, err := asj.AsynJobService.StartJobImpl(n)

	if err != nil {
		log.Println("Error while creating job ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (asj *AsyncJobController) PauseJob(ctx *gin.Context) {
	taskID := ctx.Params.ByName("task_id")

}

func (asj *AsyncJobController) ResumeJob(ctx *gin.Context) {

}

func (asj *AsyncJobController) TerminateJob(ctx *gin.Context) {

}

func (asj *AsyncJobController) StatusJob(ctx *gin.Context) {

}
