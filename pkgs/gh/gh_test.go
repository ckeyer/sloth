package gh

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/github"
	"gopkg.in/check.v1"
)

const (
	ghUser  = "test-robot"
	ghRepos = "test-project"
	ghSHA   = "58ad37eb1e8f8a97aacd42206e826aaed4512959"
)

var (
	ghToken = "fa3d9,abbb,4847,ee59f939adf4cdada538,c503127"
)

type GhSuite struct {
	ghCli *Client
}

func init() {
	ghToken = strings.Replace(ghToken, ",", "", -1)
	check.Suite(&GhSuite{
		ghCli: NewClientByToken(ghUser, ghToken),
	})
}

func Test(t *testing.T) {
	check.TestingT(t)
}

func (g *GhSuite) checkSkip(c *check.C) {
	ok, _ := strconv.ParseBool(os.Getenv("TEST_ALL"))
	if !ok {
		c.Skip("skip this.")
	}
}

func (g *GhSuite) TestListGithubRepositories(c *check.C) {
	g.checkSkip(c)

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{},
	}
	g.ghCli.Users.ListInvitations(context.Background(), nil)
	repos, resp, err := g.ghCli.Repositories.List(context.Background(), ghUser, opt)

	if err != nil {
		c.Error(err)
		if resp != nil {
			g.readBody(c, resp.Body)
		}
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

func (g *GhSuite) TestCommitPatch(c *check.C) {
	g.checkSkip(c)

	rc, resp, err := g.ghCli.Repositories.GetCommit(context.Background(), ghUser, ghRepos, ghSHA)
	if err != nil {
		c.Error(err)
		if resp != nil {
			g.readBody(c, resp.Body)
		}
		return
	}

	for _, f := range rc.Files {
		c.Logf("file: %s", mustString(f.Filename))
		oldl, newl, err := parseFilePatch(mustString(f.Patch))
		if err != nil {
			c.Error(err)
			return
		}
		c.Logf("old: %+v, new: %+v", oldl, newl)
	}
}

func (g *GhSuite) readBody(c *check.C, body io.Reader) {
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	c.Logf("readBody: %s", string(bs))
}
