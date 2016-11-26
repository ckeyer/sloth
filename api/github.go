package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/git"
)

func GithubWebhooks(rw http.ResponseWriter, req *http.Request) {
	evt := req.Header.Get("X-GitHub-Event")
	if evt == "" {
		log.Warning("unknown event type from github's webhooks")
		return
	}

	git.GetEvent(evt, req.Body)
}
