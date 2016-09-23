package api

import (
	"net/http"

	"github.com/ckeyer/sloth/git"
	// libgithub "github.com/google/go-github/github"
)

func GithubWebhooks(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	evt := ctx.req.Header.Get("X-GitHub-Event")
	if evt == "" {
		log.Warning("unknown event type from github's webhooks")
		return
	}

	git.GetEvent(evt, req.Body)
}
