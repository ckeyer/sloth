package gh

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Client struct {
	*github.Client

	UserName string
}

func NewClientByToken(user, token string) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return &Client{
		Client:   github.NewClient(tc),
		UserName: user,
	}
}
