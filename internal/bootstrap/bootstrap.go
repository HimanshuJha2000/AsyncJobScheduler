package bootstrap

import (
	"fmt"
	"github.com/DevtronLabs/headoutProj/internal"
	"github.com/DevtronLabs/headoutProj/internal/TaskRunner/model"
	"github.com/DevtronLabs/headoutProj/internal/utils"
	"github.com/tylerb/graceful"
	"log"
	"time"
)

func BaseInitAsyncJobScheduler() {
	log.Println("Async Job runner service started ...")

	model.InitInMemMap()
	router := internal.SetupRouter()

	err := graceful.RunWithErr(GetListenAddress(), utils.GracefulTimeoutDuration*time.Second, router)
	if err != nil {
		log.Println("Error occurred while starting CatPicHub server ", err)
		panic("Stopping server!!!")
	}
}

// GetListenAddress will give the address in string to listen to
func GetListenAddress() string {
	return fmt.Sprintf("%s:%d", "127.0.0.1", 8080)
}
