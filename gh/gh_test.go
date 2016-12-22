package gh

import (
	"testing"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/check.v1"
)

const (
	ghToken = "34dea836bc3e282b43a2ba774fb6046bdcfc22b4"
	ghUser  = "test-robot"
)

type ApiSuite struct {
	ghCli *github.Client
}

func init() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	check.Suite(&ApiSuite{
		ghCli: github.NewClient(tc),
	})
}

func Test(t *testing.T) {
	check.TestingT(t)
}

func (a *ApiSuite) TestMustString(c *check.C) {
	str := "abc"
	num := 123
	var nilPtr *int
	var nilStr *string
	for k, v := range map[interface{}]string{
		"":     "",
		nil:    "",
		nilPtr: "",
		nilStr: "",
		"a":    "a",
		str:    "abc",
		&str:   "abc",
		num:    "123",
		&num:   "123",
	} {
		if ret := mustString(k); v != ret {
			c.Errorf("%#v should be: <%s>, but <%s>", k, v, ret)
		}
	}
}
