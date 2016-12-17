package api

import (
	"net/http"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/gh"
	"github.com/gin-gonic/gin"
)

type GithubApp struct {
	ClientId        string
	ClientSecret    string
	AuthCallbackURL string
}

func GithubWebhooks(rw http.ResponseWriter, req *http.Request) {
	evt := req.Header.Get("X-GitHub-Event")
	if evt == "" {
		log.Warning("unknown event type from github's webhooks")
		return
	}

	gh.GetEvent(evt, req.Body)
}

// GET /github/access_url
func GetAccessURL(ctx *gin.Context) {
	ghApp := ctx.MustGet(CtxGithubApp).(GithubApp)
	u, _ := url.Parse("https://github.com/login/oauth/authorize")

	query := u.Query()
	query.Set("client_id", ghApp.ClientId)
	query.Set("redirect_uri", ghApp.AuthCallbackURL)
	scopes := []string{
		"user:email",
		"public_repo",
		"repo",
		"repo_deployment",
		"repo:status",
		"admin:repo_hook",
		"admin:org_hook",
		"admin:org",
		// "admin:public_key",
	}
	query.Add("scope", strings.Join(scopes, ","))

	u.RawQuery = query.Encode()

	GinMessage(ctx, 200, u.String())
}

// POST /github/auth
func GHAuthCallback(ctx *gin.Context) {

}
