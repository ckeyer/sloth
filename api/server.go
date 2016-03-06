package api

import (
	stdlog "log"
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/ckeyer/go-ci/lib"
	"github.com/ckeyer/go-ci/version"
	"github.com/ckeyer/go-ci/views"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

const (
	API_PREFIX = "/api"
	WEB_HOOKS  = "/webhooks"
)

var (
	log     = lib.GetLogger()
	headers map[string]string
)

type RequestContext struct {
	req    *http.Request
	render render.Render
	res    http.ResponseWriter
	params martini.Params
	mc     martini.Context
}

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
	// m.Use(httpLogger)
	m.Use(requestContext())

	m.Group(WEB_HOOKS, func(r martini.Router) {
		r.Post("/github", GithubWebhooks)
	}, MWHello)
	m.Group(API_PREFIX, func(r martini.Router) {
		r.Get("/hello", Hello)
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

func httpLogger(rw http.ResponseWriter, req *http.Request, c martini.Context) {
	log.Infof(" %s %s", req.URL.Path, req.URL.Query())
	c.Next()
	log.Infof("%s", "that is gone")
}

func requestContext() martini.Handler {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request, rnd render.Render) {
		res.Header().Set("GOCI-Version", version.GetVersion())
		res.Header().Set("X-Frame-Options", "SAMEORIGIN")
		res.Header().Set("X-XSS-Protection", "1; mode=block")

		if API_PREFIX != "" || WEB_HOOKS != "" {
			if !strings.HasPrefix(req.URL.Path, API_PREFIX) &&
				!strings.HasPrefix(req.URL.Path, WEB_HOOKS) {
				return
			}
		}

		res.Header().Set("Cache-Control", "no-cache")
		ctx := &RequestContext{
			req:    req,
			res:    res,
			render: rnd,
			params: make(map[string]string),
		}
		c.Map(ctx)

		req.ParseForm()
		if len(req.Form) > 0 {
			for k, v := range req.Form {
				ctx.params[k] = v[0]
			}
		}
	}
}
