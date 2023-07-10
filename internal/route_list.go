package internal

import (
	"github.com/DevtronLabs/headoutProj/internal/TaskRunner/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	grp := r.Group("/task")
	{
		grp.POST("/create/:sleep_time", controller.AsyncJob.StartJob)
		grp.PATCH("/pause/:task_id", controller.AsyncJob.PauseJob)
		grp.PATCH("/resume/:task_id", controller.AsyncJob.ResumeJob)
		grp.PATCH("/terminate/:task_id", controller.AsyncJob.TerminateJob)
		grp.GET("/status/:task_id", controller.AsyncJob.StatusJob)
	}
	return r
}
