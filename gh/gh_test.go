package gh

import (
	"testing"

	"github.com/google/go-github/github"
	"gopkg.in/check.v1"
)

const (
	ghToken = "9bb0bb6547ce5df09b2e6e435cff24ae1e329273"
	ghUser  = "test-robot"
)

type GhSuite struct {
	ghCli *Client
}

func init() {
	check.Suite(&GhSuite{
		ghCli: NewClientByToken(ghUser, ghToken),
	})
}

func Test(t *testing.T) {
	check.TestingT(t)
}

func (g *GhSuite) TestListGithubRepositories(c *check.C) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{},
	}
	g.ghCli.Users.ListInvitations()
	repos, resp, err := g.ghCli.Repositories.List(ghUser, opt)
	if err != nil {
		c.Error(err)
		return
	}

	if resp.StatusCode > 300 {
		c.Error(resp.Status)
		return
	}

	c.Logf("repositories: %v", len(repos))

	if len(repos) < 1 || len(repos) > 10 {
		c.Errorf("repositories count failed. count: %v", len(repos))
		return
	}

	c.Logf("repos 1: %+v", repos[0])
}
