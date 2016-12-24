package api

import (
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/admin"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func NotFound(rw http.ResponseWriter, req *http.Request) {
	for _, pre := range []string{API_PREFIX, WEB_HOOKS} {
		if strings.HasPrefix(req.URL.Path, pre) {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

// /webhooks/...
func WebhookRouter(r *gin.RouterGroup) {
	r.POST("/github", GinH(MWAuthGithubServer), GinH(GithubWebhooks))
}

func apiRouter(r *gin.RouterGroup) {
	r.GET("/_ping", ping)
	r.GET("/status", GetStatus)
	r.POST("/login", Login)
	r.POST("/signup", Registry)

	/// require login.
	func(r *gin.RouterGroup) {
		r.GET("/user/:id", GetUser)
		r.DELETE("/logout", Logout)
	}(r.Group("", MWRequireLogin))

	/// require admin.
	func(r *gin.RouterGroup) {
		r.GET("/settings", GetSettings)
		r.GET("/settings/:key", GetSettings)
		r.POST("/settings", AddSettings)
	}(r.Group("", MWRequireLogin, MWRequireAdmin))

	/// /github/...
	func(r *gin.RouterGroup) {
		r.POST("/", TODO)
		r.POST("/auth", GHAuthCallback)
		r.POST("/bind", GHBindCallback)
		r.GET("/access_url", GetAccessURL)
		r.GET("/bind_url", GetBindURL)
	}(r.Group("/github", MWLoadGithubApp))
}

func ping(ctx *gin.Context) {
	GinMessage(ctx, 200, "hi")
}

func GetStatus(ctx *gin.Context) {
	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	ret, err := admin.Status(db)
	if err != nil {
		log.Error("get status failed", err)
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, ret)
}

func TODO(ctx *gin.Context) {
	bs, _ := ioutil.ReadAll(ctx.Request.Body)

	log.WithFields(log.Fields{
		"Method":  ctx.Request.Method,
		"Path":    ctx.Request.URL.Path,
		"Remote":  ctx.Request.RemoteAddr,
		"Headers": ctx.Request.Header,
		"Agent":   ctx.Request.UserAgent(),
	}).Debug(string(bs))
	GinMessage(
		ctx,
		503,
		"Function is under development...",
		ctx.Request.Method,
		ctx.Request.URL.Path,
	)
}
