package git

import (
	"github.com/ckeyer/go-ci/lib"
)

var log = lib.GetLogger()

type Giter interface {
	Clone(repo, ref string) (string, error)
}
