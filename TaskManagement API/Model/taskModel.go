package Model

import (
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      WorkStatus `json:"status"`
	Deadline    time.Time  `json:"deadline"`
}
