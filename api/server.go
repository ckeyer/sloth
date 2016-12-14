package api

import (
	stdlog "log"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/version"
	"github.com/ckeyer/sloth/views"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

const (
	API_PREFIX = "/api"
	WEB_HOOKS  = "/webhooks"
)

var headHandle = cors.New(cors.Config{
	AbortOnError:    false,
	AllowAllOrigins: true,
	// AllowedOrigins:   []string{"*"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	AllowedHeaders:   []string{"Content-Type", "Limt", "Offset", "Origin", "Accept"},
	ExposedHeaders:   []string{"Record-Count", "Limt", "Offset", "Content-Type"},
	AllowCredentials: true,
	MaxAge:           24 * time.Hour,
})

func Serve(listenAddr string, db *mgo.Database) {

	gr := NewGin()

	gr.NoRoute(GinH(views.New()))

	gr.Use(headHandle)

	store := sessions.NewCookieStore([]byte("secret"))
	gr.Use(sessions.Sessions("ck-session", store))

	// m.Use(sessions.Sessions("sloth", utils.GetCookieStore()))
	gr.Use(requestContext())

	gr.Use(GinLogger)

	// routers...
	// /webhooks
	WebhookRouter(gr.Group(WEB_HOOKS))
	// /api
	apiRouter(gr.Group(API_PREFIX))

	logger := log.StandardLogger()
	server := &http.Server{
		Handler:  gr,
		Addr:     listenAddr,
		ErrorLog: stdlog.New(logger.Writer(), "", 0),
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

		if API_PREFIX != "" || WEB_HOOKS != "" {
			if !strings.HasPrefix(req.URL.Path, API_PREFIX) &&
				!strings.HasPrefix(req.URL.Path, WEB_HOOKS) {
				return
			}
		}
		res.Header().Set("Cache-Control", "no-cache")

		/// TODO: set & load session
		ss := sessions.Default(ctx)
		_ = ss.Get("key")
		// ctx.Set("ss", store.Get(ctx.Request, "UserToken"))
	}
}
