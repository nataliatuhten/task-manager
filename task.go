package model

import (
	"time"
)

type Task struct {
	ID           int
	Title        string
	Description  string
	ProjectID    int    // ID of the project the task belongs to
	AssigneeToID int    // ID of the user responsible for the task
	Status       string // e.g: to do, in progress, done, review...
	DueDate      time.Time
	CreatedAt    time.Time
}
