package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

func WebhookRouter(r *gin.RouterGroup) {
	r.POST("/github", GinH(MWAuthGithubServer), GinH(GithubWebhooks))
}

func apiRouter(r *gin.RouterGroup) {
	r.GET("/_ping", ping)
	r.POST("/login", TODO)
	r.POST("/signup", TODO)

	func(r *gin.RouterGroup) {
		r.DELETE("/logout", TODO)
	}(r.Group("/user"))
}

func ping(ctx *gin.Context) {
	GinMessage(ctx, 200, "hi")
}
