package main

import (
	"fmt"
	"github.com/DevtronLabs/headoutProj/internal/bootstrap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	ID           string
	Status       string
	Cursor       int
	Target       int
	IsPaused     bool
	IsTerminated bool
	mu           sync.Mutex
}

var tasks map[string]*Task

func main() {
	tasks = make(map[string]*Task)

	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/pause", pauseHandler)
	http.HandleFunc("/resume", resumeHandler)
	http.HandleFunc("/terminate", terminateHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		http.Error(w, "Invalid value for 'n'", http.StatusBadRequest)
		return
	}

	task := &Task{
		ID:           generateID(),
		Status:       "RUNNING",
		Cursor:       0,
		Target:       n,
		IsPaused:     false,
		IsTerminated: false,
	}

	go runTask(task)

	tasks[task.ID] = task

	fmt.Fprintf(w, `{"statuscode": 200, "id": "%s"}`, task.ID)
}

func pauseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task.mu.Lock()
	defer task.mu.Unlock()

	task.IsPaused = true
	task.Status = "PAUSED"

	fmt.Fprintf(w, `{"statuscode": 200}`)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task.mu.Lock()
	defer task.mu.Unlock()

	task.IsPaused = false
	task.Status = "RUNNING"

	fmt.Fprintf(w, `{"statuscode": 200}`)
}

func terminateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task.mu.Lock()
	defer task.mu.Unlock()

	task.IsTerminated = true
	task.Status = "TERMINATED"

	fmt.Fprintf(w, `{"statuscode": 200}`)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task.mu.Lock()
	defer task.mu.Unlock()

	fmt.Fprintf(w, `{"statuscode": 200, "jobStatus": "%s", "cursor": %d}`, task.Status, task.Cursor)
}

func runTask(task *Task) {
	for task.Cursor < task.Target {
		task.mu.Lock()
		if task.IsTerminated {
			task.mu.Unlock()
			return
		}
		if task.IsPaused {
			task.mu.Unlock()
			time.Sleep(time.Second)
			continue
		}
		task.Cursor++
		task.mu.Unlock()

		// Simulating the task by sleeping for 1 second
		time.Sleep(time.Second)
	}

	task.mu.Lock()
	task.Status = "COMPLETED"
	task.mu.Unlock()
}

func generateID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36)
}

func main() {
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
