package shell

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	// github's url
	GithubURL = "https://github.com/"
	WorkDir   = "/tmp/go-ci/src/github.com/"
)

// ShellGit
type ShellGit struct {
}

// Clone
func (s *ShellGit) Clone(fullName, ref string) (string, error) {
	path := WorkDir + fullName
	if err := s.initDirPath(path); err != nil {
		return "", err
	}

	repo := GithubURL + fullName
	cmd := exec.Command("git", "clone", repo, "-b", ref, path)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Shell clone exec %s, %s", repo, err)
	}
	return path, nil
}

func (s *ShellGit) initDirPath(path string) (err error) {
	if err = os.RemoveAll(path); err != nil {
		return fmt.Errorf("Initialization clone: rm -f %s, %s", path, err)
	}

	if err = os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("Initialization clone: mkdir %s, %s", path, err)
	}
	return
}
