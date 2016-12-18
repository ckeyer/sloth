package gh

import (
	"encoding/json"
	"io"

	log "github.com/Sirupsen/logrus"
	libgithub "github.com/google/go-github/github"
)

const (
	EVT_PULL_REQUEST   = "pull_request"
	EVT_COMMIT_COMMENT = "commit_comment"
)

func GetEvent(evt string, data io.Reader) {
	var ret interface{}
	switch evt {
	case EVT_PULL_REQUEST:
		ret = new(libgithub.PullRequestEvent)
	case EVT_COMMIT_COMMENT:
		ret = new(libgithub.CommitCommentEvent)
	default:
		return
	}
	err := json.NewDecoder(data).Decode(ret)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug(ret)
}

func try() {
	cli := libgithub.NewClient(nil)
	cli.NewRequest(method, urlStr, body)
}
