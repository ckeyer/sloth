package version

import (
	"strings"
)

var (
	version, gitCommit string
)

func GetVersion() string {
	return version
}

func GetGitCommit() string {
	return gitCommit
}

func GetCompleteVersion() string {
	return strings.Join([]string{GetVersion(), GetGitCommit()}, "-")
}
