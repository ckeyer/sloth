package job

import (
	"time"
)

// Job:
type Job interface {
	Run() error
	Kill()
}

// JobManager:
type JobManager interface {
	Add(Job) string
	Del(string)
	GetStatus(int) *Status
}

// Status:
type Status struct {
	Status    string    `json:"status"`
	Err       string    `json:"err,omitempty"`
	Out       []byte    `json:"out"`
	StartTime time.Time `json:"start_time"`
	OverTime  time.Time `json:"over_time"`
}
