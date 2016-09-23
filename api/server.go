package api

import (
	stdlog "log"
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/lib"
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

var (
	log     = lib.GetLogger()
	headers map[string]string
)

func init() {
	headers = make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "GET,OPTIONS,POST,DELETE"
	headers["Access-Control-Allow-Credentials"] = "true"
	headers["Access-Control-Max-Age"] = "864000"
	headers["Access-Control-Expose-Headers"] = "Record-Count,Limt,Offset,Content-Type"
	headers["Access-Control-Allow-Headers"] = "Limt,Offset,Content-Type,Origin,Accept,Authorization"
}

func Serve(listenAddr string) {
	m := NewMartini()
	view := views.New()

	m.NotFound(NotFound, view.ServeHTTP)

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     strings.Split(headers["Access-Control-Allow-Origin"], ","),
		AllowMethods:     strings.Split(headers["Access-Control-Allow-Methods"], ","),
		AllowHeaders:     strings.Split(headers["Access-Control-Allow-Headers"], ","),
		ExposeHeaders:    strings.Split(headers["Access-Control-Expose-Headers"], ","),
		AllowCredentials: true,
		MaxAge:           time.Second * 864000,
	}))

	m.Use(sessions.Sessions("kyt-api", lib.GetCookieStore()))
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

func router() {

}
