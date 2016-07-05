package job

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os/exec"
	"time"
)

type ShellJob struct {
	ID  string
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
	_, err := s.Cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (s *ShellJob) Kill() {

}

func (s *ShellJob) GetStatus() *Status {
	return nil
}
