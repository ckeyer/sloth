package api

import (
	logpkg "log"
	"net/http"
	"strings"

	"github.com/ckeyer/sloth/api/views"
	"github.com/ckeyer/sloth/version"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	// royter prefix
	API_PREFIX = "/api"
	WEB_HOOKS  = "/webhooks"

	// gin.Context.Set key
	CtxKeyMgoDB     = "mgodb"
	CtxKeyUserAuth  = "userauth"
	CtxKeyUser      = "user"
	CtxKeyGithubApp = "githubapp"
	CtxKeyGithubCli = "githubclient"
)

// Serve: main serve.
func Serve(listenAddr string, db *mgo.Database) {

	gr := NewGin()

	gr.NoRoute(GinH(views.New()))

	gr.Use(CorsHandle)

	store := sessions.NewCookieStore([]byte("secret"))
	gr.Use(sessions.Sessions("ck-session", store))

	gr.Use(requestContext())

	gr.Use(GinLogger)
	gr.Use(SetMgoDB(db))

	// routers...
	// /webhooks
	WebhookRouter(gr.Group(WEB_HOOKS))
	// /api
	apiRouter(gr.Group(API_PREFIX))

	logger := log.StandardLogger()
	server := &http.Server{
		Handler:  gr,
		Addr:     listenAddr,
		ErrorLog: logpkg.New(logger.Writer(), "", 0),
	}

	log.Info("server is starting on ", listenAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
	}
}

func requestContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, req := ctx.Writer, ctx.Request

		res.Header().Set("Sloth-Version", version.GetVersion())
		res.Header().Set("X-XSS-Protection", "1; mode=block")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		if API_PREFIX != "" || WEB_HOOKS != "" {
			if !strings.HasPrefix(req.URL.Path, API_PREFIX) &&
				!strings.HasPrefix(req.URL.Path, WEB_HOOKS) {
				return
			}
		}

		res.Header().Set("Cache-Control", "no-cache")
	}
}
