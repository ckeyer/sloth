package job

import (
	"crypto/sha1"
	"fmt"
	"sync"
	"time"
)

type JobManage struct {
	sync.Mutex
	Jobs map[string]Job
}

func (jm *JobManage) Add(job Job) string {
	var jobID string
	for {
		jobID = fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprint(time.Now().UnixNano()))))
		if _, exists := jm.Jobs[jobID]; !exists {
			break
		}
	}
	jm.Jobs[jobID] = job
	return jobID
}

func (jm *JobManage) Del(job_id string) {
	if job, exi := jm.Jobs[job_id]; exi {
		job.Kill()
		delete(jm.Jobs, job_id)
	}
}
