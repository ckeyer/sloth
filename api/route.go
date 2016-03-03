package api

import (
	"github.com/go-martini/martini"
	"net/http"
)

type routerHandle func(martini.Router)

func APIRouter() routerHandle {
	return func(r martini.Router) {
		r.Get("/hello", Hello)
	}
}

func Hello(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	ctx.render.Text(http.StatusOK, "hello, "+ctx.req.RemoteAddr)
}

func NotFound(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
}
