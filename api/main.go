package main

import (
	"flag"
	"github.com/DevtronLabs/headoutProj/internal/bootstrap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	done := make(chan struct{})

	// Wait for termination signal (SIGINT, SIGTERM, SIGKILL)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalChannel
		close(done)
	}()

	bootstrap.BaseInitAsyncJobScheduler()

	<-done
}
