package git

import (
	"github.com/ckeyer/sloth/utils"
)

var log = lib.GetLogger()

type Giter interface {
	Clone(repo, ref string) (string, error)
}
