package checks

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// GetDirAndFiles return dir, filenames, error
func GetDirAndFiles(args []string) (string, []string, error) {
	dir := ""
	files := []string{}

	for _, arg := range args {
		fi, err := os.Stat(arg)
		if err != nil {
			log.WithFields(log.Fields{
				"path":  arg,
				"error": err,
			}).Error("...")
			return "", nil, err
		}

		if fi.IsDir() {
			if dir != "" {
				return "", nil, fmt.Errorf("not support multiple directory")
			}
			dir = arg
		} else {
			files = append(files, arg)
		}
	}

	if dir == "" {
		dir = "."
	}
	if len(files) == 0 {
		var err error
		files, _, err = GoFiles(dir)
		if err != nil {
			return "", nil, err
		}
	}

	return dir, files, nil
}
