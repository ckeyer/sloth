package api

import (
	"github.com/ckeyer/sloth/version"
	"github.com/gin-gonic/gin"
	stdlog "log"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/views"
	"github.com/martini-contrib/cors"
)

const (
	API_PREFIX = "/api"
	WEB_HOOKS  = "/webhooks"
)

var headHandle = cors.Allow(&cors.Options{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "OPTIONS", "POST", "DELETE"},
	AllowHeaders:     []string{"Limt,Offset,Content-Type,Origin,Accept,Authorization"},
	ExposeHeaders:    []string{"Record-Count", "Limt", "Offset", "Content-Type"},
	AllowCredentials: true,
	MaxAge:           time.Second * 864000,
})

func Serve(listenAddr string) {
	gr := NewGin()

	gr.NoRoute(GinH(views.New()))

	gr.Use(GinH(headHandle))

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
	}
}
