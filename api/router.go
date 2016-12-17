package api

import (
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/admin"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func TODO(ctx *gin.Context) {
	GinMessage(
		ctx,
		503,
		"Function is under development...",
		ctx.Request.Method,
		ctx.Request.URL.Path,
	)
}

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
		r.GET("/access_url", GetAccessURL)
	}(r.Group("/github", MWLoadGithubApp))
}

func ping(ctx *gin.Context) {
	GinMessage(ctx, 200, "hi")
}

func GetStatus(ctx *gin.Context) {
	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)
	ret, err := admin.Status(db)
	if err != nil {
		log.Error("get status failed", err)
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, ret)
}
