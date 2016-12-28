package job

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os/exec"
	"time"
)

// ShellJob:
type ShellJob struct {
	ID  string
	Cmd *exec.Cmd
	out *bytes.Buffer
}

// NewShellJob
func NewShellJob(cmd *exec.Cmd) *ShellJob {
	bs := fmt.Sprintf("%x\n", sha1.Sum([]byte(fmt.Sprint(time.Now().UnixNano()))))
	return &ShellJob{
		ID:  bs,
		Cmd: cmd,
		out: new(bytes.Buffer),
	}
}

// GetID
func (s *ShellJob) GetID() string {
	return s.ID
}

// Run
func (s *ShellJob) Run() error {
	_, err := s.Cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

// Kill
func (s *ShellJob) Kill() {

}

// GetStatus
func (s *ShellJob) GetStatus() *Status {
	return nil
}
