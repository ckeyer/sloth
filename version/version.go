package version

import (
	"strings"
)

var (
	version, gitCommit string
)

// GetVersion
func GetVersion() string {
	return version
}

// GetGitCommit
func GetGitCommit() string {
	return gitCommit
}

// GetCompleteVersion
func GetCompleteVersion() string {
	return strings.Join([]string{GetVersion(), GetGitCommit()}, "-")
}
