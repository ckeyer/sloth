package shell

import (
	"testing"
)

func Test(t *testing.T) {
	git := new(ShellGit)
	git.Clone("ckeyer/t", "master")
}
