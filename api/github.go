package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/gh"
	"github.com/gin-gonic/gin"
)

type GithubApp struct {
	ClientID        string
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
	ghApp := ctx.MustGet(CtxGithubApp).(*GithubApp)
	u, _ := url.Parse("https://github.com/login/oauth/authorize")

	query := u.Query()
	query.Set("client_id", ghApp.ClientID)
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
	ghApp := ctx.MustGet(CtxGithubApp).(*GithubApp)
	code := ctx.Query("code")

	u, _ := url.Parse("https://github.com/login/oauth/access_token")

	query := u.Query()
	query.Set("client_id", ghApp.ClientID)
	query.Set("redirect_uri", ghApp.AuthCallbackURL)
	query.Set("client_secret", ghApp.ClientSecret)
	query.Set("code", code)
	u.RawQuery = query.Encode()

	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("get github-app token failed, %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	var ghToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	err = json.NewDecoder(resp.Body).Decode(&ghToken)
	if err != nil {
		log.Errorf("read github-app token failed, %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	log.Debugf("get github-app token: %+v", ghToken)
	ctx.Redirect(302, "/")
}
