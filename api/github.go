package api

import (
	"net/http"

	"github.com/ckeyer/go-ci/github"
	// libgithub "github.com/google/go-github/github"
)

func GithubWebhooks(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	evt := ctx.req.Header.Get("X-GitHub-Event")
	if evt == "" {
		log.Warning("unknown event type from github's webhooks")
		return
	}

	github.GetEvent(evt, req.Body)
}
