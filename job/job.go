package job

import (
	"time"
)

type Job interface {
	Run() error
	Kill()
}

type JobManager interface {
	Add(Job) string
	Del(string)
	GetStatus(int) *Status
}

type Status struct {
	Status    string    `json:"status"`
	Err       string    `json:"status,omitempty"`
	Out       []byte    `json:"out"`
	StartTime time.Time `json:"start_time"`
	OverTime  time.Time `json:"over_time"`
}
