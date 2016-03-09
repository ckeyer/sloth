package api

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

type RequestContext struct {
	req    *http.Request
	render render.Render
	res    http.ResponseWriter
	params martini.Params
	ss     sessions.Session
	mc     martini.Context
	u      *types.User
}

func (r *RequestContext) Error() {

}

func requestContext() martini.Handler {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request, rnd render.Render, ss sessions.Session) {
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

		uid, ok := ss.Get("user").(int64)
		if !ok || uid == 0 {
			return
		}
		ctx := &RequestContext{
			req:    req,
			ss:     ss,
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
