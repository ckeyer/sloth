package api

import (
	"net/http"
	"strings"
)

func Hello(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	ctx.render.Text(http.StatusOK, "hello fuck, "+ctx.req.RemoteAddr)
}

func NotFound(rw http.ResponseWriter, req *http.Request) {
	for _, pre := range []string{API_PREFIX, WEB_HOOKS} {
		if strings.HasPrefix(req.URL.Path, pre) {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
