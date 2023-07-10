package model

import "sync"

type Task struct {
	ID           string
	Status       string
	Cursor       int
	Target       int
	IsPaused     bool
	IsTerminated bool
	Mu           sync.Mutex
}

var tasksMap map[string]*Task

func InitInMemMap() map[string]*Task {
	tasksMap = make(map[string]*Task)
	return tasksMap
}

func AddToInMem(task *Task) {
	tasksMap[task.ID] = task
}

func GetMemMap() map[string]*Task {
	return tasksMap
}
