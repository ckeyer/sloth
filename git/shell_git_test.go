package git

import (
	"testing"
)

func Test(t *testing.T) {
	git := new(ShellGit)
	git.Clone("ckeyer/kyt-api", "master")
}
