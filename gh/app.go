package gh

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/types"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type App struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

func (a *App) AuthURL() string {
	return strings.TrimSuffix(a.CallbackURL, "/") + "/auth"
}

func (a *App) BindURL() string {
	return strings.TrimSuffix(a.CallbackURL, "/") + "/bind"
}

// bind or auth
func (a *App) AccessURL(use string) *url.URL {
	u, _ := url.Parse("https://github.com/login/oauth/authorize")
	callback := a.AuthURL()
	if use == "bind" {
		callback = a.BindURL()
	}
	query := u.Query()
	query.Set("client_id", a.ClientID)
	query.Set("redirect_uri", callback)
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

	return u
}

func (a *App) GetToken(code string) (string, error) {
	u, _ := url.Parse("https://github.com/login/oauth/access_token")

	query := u.Query()
	query.Set("client_id", a.ClientID)
	query.Set("client_secret", a.ClientSecret)
	query.Set("code", code)
	u.RawQuery = query.Encode()

	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("get github-app token failed, %s", err.Error())
		return "", err
	}

	var ghToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	err = json.NewDecoder(resp.Body).Decode(&ghToken)
	if err != nil {
		log.Errorf("read github-app token failed, %s", err.Error())
		return "", err
	}
	log.Debugf("get github-app token: %+v", ghToken)
	return ghToken.AccessToken, nil
}

func (a *App) GetUserAccount(token string) (*types.GithubAccount, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	user, _, err := client.Users.Get("")
	if err != nil {
		log.Errorf("GetUserAccount: %s", err)
		return nil, err
	}
	return &types.GithubAccount{
		ID:       *user.ID,
		Login:    mustString(user.Login),
		Name:     mustString(user.Name),
		Email:    mustString(user.Email),
		Location: mustString(user.Location),
		Type:     mustString(user.Type),

		Token: types.Password(token),
	}, nil
}
