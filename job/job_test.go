package job

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type JobSuite struct{}

func init() {
	check.Suite(new(JobSuite))
}

func (j *JobSuite) TestNewID(c *check.C) {
	c.Check(NewID(), check.Equals, 1)
	c.Check(NewID(), check.Equals, 2)
	c.Check(NewID(), check.Equals, 3)
	c.Check(NewID(), check.Equals, 4)
}
