package git

import (
	"github.com/ckeyer/sloth/lib"
)

var log = lib.GetLogger()

type Giter interface {
	Clone(repo, ref string) (string, error)
}
