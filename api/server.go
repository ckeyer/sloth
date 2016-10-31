package api

import (
	stdlog "log"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/utils"
	"github.com/ckeyer/sloth/views"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

const (
	API_PREFIX = "/api"
	WEB_HOOKS  = "/webhooks"
)

func Serve(listenAddr string) {
	m := NewMartini()
	view := views.New()

	m.NotFound(NotFound, view.ServeHTTP)

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "OPTIONS", "POST", "DELETE"},
		AllowHeaders:     []string{"Limt,Offset,Content-Type,Origin,Accept,Authorization"},
		ExposeHeaders:    []string{"Record-Count", "Limt", "Offset", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           time.Second * 864000,
	}))

	m.Use(sessions.Sessions("sloth", lib.GetCookieStore()))
	m.Use(requestContext())

	m.Group(WEB_HOOKS, func(r martini.Router) {
		r.Post("/github", WMAuthGithubServer, GithubWebhooks)
	}, MWHello)

	m.Group(API_PREFIX, func(r martini.Router) {
		r.Get("/hello", Hello)
		r.Post("/login", Login)
	}, MWHello)

	logger := logrus.StandardLogger()
	server := &http.Server{
		Handler:  m,
		Addr:     listenAddr,
		ErrorLog: stdlog.New(logger.Writer(), "", 0),
	}

	log.Notice("server is starting on ", listenAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
	}
}

func NewMartini() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{Martini: m, Router: r}
}
