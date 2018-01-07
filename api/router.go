package api

import (
	"github.com/gin-gonic/gin"
)

// WebhookRouter: /webhooks/...
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

		// projects api
		func(r *gin.RouterGroup) {
			r.GET("", ListProjects)
			r.POST("", NewProject)
			r.GET("/:id", GetProject)
			r.POST("/:id", UpdateProject)
			r.DELETE("/:id", RemoveProject)
		}(r.Group("/project"))
	}(r.Group("", MWRequireLogin))

	/// system configure, require admin.
	func(r *gin.RouterGroup) {
		r.GET("/settings", GetSettings)
		r.GET("/settings/:key", GetSettings)
		r.POST("/settings", AddSettings)
	}(r.Group("", MWRequireLogin, MWRequireAdmin))

	/// Github Apps /github/...  public.
	func(r *gin.RouterGroup) {
		r.POST("/", TODO)
		r.POST("/auth", GHAuthCallback)
		r.POST("/bind", GHBindCallback)
		r.GET("/access_url", GetAccessURL)
		r.GET("/bind_url", GetBindURL)
	}(r.Group("/github", MWLoadGithubApp))
}
