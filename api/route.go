package api

import (
	"net/http"
	"strings"
)

func Hello(rw http.ResponseWriter, req *http.Request) {

}

func NotFound(rw http.ResponseWriter, req *http.Request) {
	for _, pre := range []string{API_PREFIX, WEB_HOOKS} {
		if strings.HasPrefix(req.URL.Path, pre) {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
