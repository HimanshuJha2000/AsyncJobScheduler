package bootstrap

import (
	"github.com/DevtronLabs/headoutProj/internal"
	"github.com/DevtronLabs/headoutProj/internal/TaskRunner/model"
	"log"
)

func BaseInitAsyncJobScheduler() {
	log.Println("Async Job runner service started ...")
	internal.SetupRouter()
	model.InitInMemMap()
}
