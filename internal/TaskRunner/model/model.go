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

func InitInMemMap() {
	tasksMap = make(map[string]*Task)
}

func AddToInMem(task *Task) {
	tasksMap[task.ID] = task
}

func GetInMemMap() map[string]*Task {
	return tasksMap
}
