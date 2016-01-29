package version

import (
	"fmt"
	"testing"

	"gopkg.in/check.v1"
)

type VersionSuite struct{}

var _ = check.Suite(&VersionSuite{})

func Test(t *testing.T) { check.TestingT(t) }

func (v *VersionSuite) TestVersion(c *check.C) {
	fmt.Println("...")
	c.Check(GetVersion(), check.Not(check.Equals), "")
}
