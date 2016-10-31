package utils

import (
	logpkg "github.com/ckeyer/go-log"
)

var log *logpkg.Logger

func init() {
	if log == nil {
		log = logpkg.GetDefaultLogger("go-ci")
	}
}

func GetLogger() *logpkg.Logger {
	return log
}
