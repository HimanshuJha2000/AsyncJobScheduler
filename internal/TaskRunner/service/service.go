package service

import (
	"fmt"
	"github.com/DevtronLabs/headoutProj/internal/TaskRunner/model"
	"github.com/DevtronLabs/headoutProj/internal/utils"
	"net/http"
	"time"
)

type Service struct{}

func (svc Service) StartJobImpl(sleepTime int) (int, map[string]interface{}, error) {
	task := &model.Task{
		ID:           utils.GenerateID(),
		Status:       utils.JOB_RUNNING,
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
		"success":         "Task created and started successfully",
	}, nil
}

func (svc Service) PauseJobImpl(taskID string) (int, map[string]interface{}, error) {

	tasks := model.GetInMemMap()
	task, ok := tasks[taskID]

	if !ok {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task doesn't exist",
		}, fmt.Errorf("error")
	}

	task.Mu.Lock()
	defer task.Mu.Unlock()

	if task.IsTerminated {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task is already terminated",
		}, fmt.Errorf("error")
	}

	if task.IsPaused {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task is already paused",
		}, fmt.Errorf("error")
	}

	task.IsPaused = true
	task.Status = utils.JOB_PAUSED

	return http.StatusOK, map[string]interface{}{
		"task_ID":     task.ID,
		"task_status": task.Status,
		"success":     "Task added successfully",
	}, nil
}

func (svc Service) ResumeJobImpl(taskID string) (int, map[string]interface{}, error) {
	tasks := model.GetInMemMap()
	task, ok := tasks[taskID]

	if !ok {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task doesn't exist",
		}, fmt.Errorf("error occured")
	}

	task.Mu.Lock()
	defer task.Mu.Unlock()

	if task.IsTerminated {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task is already terminated",
		}, fmt.Errorf("error")
	}

	if task.IsPaused {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task is already paused",
		}, fmt.Errorf("error")
	}

	task.IsPaused = false
	task.Status = utils.JOB_RUNNING

	return http.StatusOK, map[string]interface{}{
		"task_ID":     task.ID,
		"task_status": task.Status,
		"success":     "Task restarted successfully",
	}, nil
}

func (svc Service) TerminateJobImpl(taskID string) (int, map[string]interface{}, error) {
	tasks := model.GetInMemMap()
	task, ok := tasks[taskID]

	if !ok {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task doesn't exist",
		}, fmt.Errorf("error occured")
	}

	task.Mu.Lock()
	defer task.Mu.Unlock()

	if task.IsTerminated {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task is already terminated",
		}, fmt.Errorf("error")
	}

	task.IsTerminated = true
	task.Status = utils.JOB_TERMINATED

	return http.StatusOK, map[string]interface{}{
		"task_ID":     task.ID,
		"task_status": task.Status,
		"success":     "Task terminated successfully",
	}, nil
}

func (svc Service) JobStatusImpl(taskID string) (int, map[string]interface{}, error) {
	tasks := model.GetInMemMap()
	task, ok := tasks[taskID]

	if !ok {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       "Task doesn't exist",
		}, fmt.Errorf("error occured")
	}

	task.Mu.Lock()
	defer task.Mu.Unlock()

	return http.StatusOK, map[string]interface{}{
		"task_ID":     task.ID,
		"task_status": task.Status,
		"success":     "Task added successfully",
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
	task.Status = utils.JOB_COMPLETED
	task.Mu.Unlock()
}
