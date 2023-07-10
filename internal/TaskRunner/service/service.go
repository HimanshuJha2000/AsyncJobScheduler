package service

import (
	"github.com/DevtronLabs/headoutProj/internal/TaskRunner/model"
	"github.com/DevtronLabs/headoutProj/internal/utils"
	"net/http"
	"time"
)

type Service struct{}

func (svc Service) StartJobImpl(sleepTime int) (int, map[string]interface{}, error) {
	task := &model.Task{
		ID:           utils.GenerateID(),
		Status:       "RUNNING",
		Cursor:       0,
		Target:       sleepTime,
		IsPaused:     false,
		IsTerminated: false,
	}

	go runTask(task)

	model.AddToInMem(task)

	return http.StatusOK, map[string]interface{}{
		"task_ID":         task.ID,
		"task_status":     task.Status,
		"task_sleep_time": task.Target,
		"success":         "Task added successfully",
	}, nil
}

func (svc Service) PauseJobImpl(taskID int) (int, map[string]interface{}, error) {

	return http.StatusOK, map[string]interface{}{
		"task_ID":         task.ID,
		"task_status":     task.Status,
		"task_sleep_time": task.Target,
		"success":         "Task added successfully",
	}, nil
}

func (svc Service) ResumeJobImpl(sleepTime int) (int, map[string]interface{}, error) {
	task := &model.Task{
		ID:           utils.GenerateID(),
		Status:       "RUNNING",
		Cursor:       0,
		Target:       sleepTime,
		IsPaused:     false,
		IsTerminated: false,
	}

	go runTask(task)

	model.AddToInMem(task)

	return http.StatusOK, map[string]interface{}{
		"task_ID":         task.ID,
		"task_status":     task.Status,
		"task_sleep_time": task.Target,
		"success":         "Task added successfully",
	}, nil
}

func (svc Service) TerminateJobImpl(sleepTime int) (int, map[string]interface{}, error) {
	task := &model.Task{
		ID:           utils.GenerateID(),
		Status:       "RUNNING",
		Cursor:       0,
		Target:       sleepTime,
		IsPaused:     false,
		IsTerminated: false,
	}

	go runTask(task)

	model.AddToInMem(task)

	return http.StatusOK, map[string]interface{}{
		"task_ID":         task.ID,
		"task_status":     task.Status,
		"task_sleep_time": task.Target,
		"success":         "Task added successfully",
	}, nil
}

func (svc Service) TerminateJobImpl(sleepTime int) (int, map[string]interface{}, error) {
	task := &model.Task{
		ID:           utils.GenerateID(),
		Status:       "RUNNING",
		Cursor:       0,
		Target:       sleepTime,
		IsPaused:     false,
		IsTerminated: false,
	}

	go runTask(task)

	model.AddToInMem(task)

	return http.StatusOK, map[string]interface{}{
		"task_ID":         task.ID,
		"task_status":     task.Status,
		"task_sleep_time": task.Target,
		"success":         "Task added successfully",
	}, nil
}

func runTask(task *model.Task) {
	for task.Cursor < task.Target {
		task.Mu.Lock()
		if task.IsTerminated {
			task.Mu.Unlock()
			return
		}
		if task.IsPaused {
			task.Mu.Unlock()
			time.Sleep(time.Second)
			continue
		}
		task.Cursor++
		task.Mu.Unlock()

		// Simulating the task by sleeping for 1 second
		time.Sleep(time.Second)
	}

	task.Mu.Lock()
	task.Status = "COMPLETED"
	task.Mu.Unlock()
}
