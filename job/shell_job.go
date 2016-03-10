package job

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os/exec"
	"time"
)

type ShellJob struct {
	ID  int
	Cmd *exec.Cmd
	out *bytes.Buffer
}

func NewShellJob(cmd *exec.Cmd) *ShellJob {
	bs := fmt.Sprintf("%x\n", sha1.Sum([]byte(fmt.Sprint(time.Now().UnixNano()))))
	return &ShellJob{
		ID:  bs,
		Cmd: cmd,
		out: new(bytes.Buffer),
	}
}

func (s *ShellJob) GetID() string {
	return s.ID
}

func (s *ShellJob) Run() error {
	out, err := s.Cmd.CombinedOutput()
	if err != nil {
		return error
	}
}

func (s *ShellJob) Kill() {

}

func (s *ShellJob) GetStatus() *Status {

}
