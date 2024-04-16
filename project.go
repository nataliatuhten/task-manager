package model

import (
	"time"
)

type Project struct {
	ID           int
	Name         string
	Description  string
	OwnerId      int    // ID of the user who created the project
	Participants []int  // ID of the users who are participating the project
	Tasks        []Task //Project Tasks
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
