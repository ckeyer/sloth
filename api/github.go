package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/admin"
	"github.com/ckeyer/sloth/gh"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

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
	ghApp := ctx.MustGet(CtxGithubApp).(*gh.App)
	u := ghApp.AccessURL("auth")
	GinMessage(ctx, 200, u.String())
}

// GET /github/bind_url
func GetBindURL(ctx *gin.Context) {
	ghApp := ctx.MustGet(CtxGithubApp).(*gh.App)
	u := ghApp.AccessURL("bind")
	GinMessage(ctx, 200, u.String())
}

// GET /github/auth
func GHAuthCallback(ctx *gin.Context) {
	ghApp := ctx.MustGet(CtxGithubApp).(*gh.App)
	code := getGHCode(ctx.Request.Body)
	if code == "" {
		GinError(ctx, 400, "invalid code")
		return
	}
	token, err := ghApp.GetToken(code)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	ghAccount, err := ghApp.GetUserAccount(token)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}
	log.Debugf("GHAuthCallback: %+v", ghAccount)

	statusCode := 200
	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)
	user, err := admin.GetUserByGHAccount(db, ghAccount.ID)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("GHAuthCallback: unknown error, %s", err)
		GinError(ctx, 500, err)
		return
	} else if err == mgo.ErrNotFound {
		// 未查到，注册新用户
		user, err = (*admin.User).RegistryByGHAccount(nil, db, ghAccount)
		if err != nil {
			log.Errorf("GHAuthCallback: registry user failed, %s", err)
			GinError(ctx, 500, err)
			return
		}
		log.Debugf("GHAuthCallback: registry a user by github account, %+v", user)
		statusCode = 201
	}

	ua := admin.NewUserAuth(user.ID, time.Now().Add(time.Hour*24*90))
	if err := ua.Insert(db); err != nil {
		GinError(ctx, 500, err)
		return
	}

	ret := map[string]interface{}{
		"user":      user,
		"user_auth": ua,
	}
	ctx.JSON(statusCode, ret)
}

// GET /github/bind
func GHBindCallback(ctx *gin.Context) {
	ghApp := ctx.MustGet(CtxGithubApp).(*gh.App)
	code := getGHCode(ctx.Request.Body)
	if code == "" {
		GinError(ctx, 400, "invalid code")
		return
	}

	token, err := ghApp.GetToken(code)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	ghAccount, err := ghApp.GetUserAccount(token)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	log.Debugf("GHBindCallback: %+v", ghAccount)
	ctx.Redirect(302, "/user")
}

// 获取body中的code
func getGHCode(r io.Reader) string {
	v := map[string]string{}
	err := json.NewDecoder(r).Decode(&v)
	if err != nil {
		log.Errorf("invalid code %s", err.Error())
		return ""
	}
	if code, ok := v["code"]; !ok {
		log.Errorf("required code")
		return ""
	} else {
		return code
	}
}
