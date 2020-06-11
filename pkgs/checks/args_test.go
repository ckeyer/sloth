package checks

import (
	"testing"
)

func TestFmtArgs(t *testing.T) {
	ArgsTable := []struct {
		args  []string
		dir   string
		files []string
		isErr bool
	}{
		{[]string{}, "", []string{}, false},
		{[]string{"args.go"}, "", []string{"args.go"}, false},
	}

	for _, v := range ArgsTable {
		dir, files, err := GetDirAndFiles(v.args)
		if err != nil {
			if v.isErr {
				continue
			}

			t.Error(err)
			return
		}
		t.Logf("return: %+v: %+v", dir, files)
	}
}
