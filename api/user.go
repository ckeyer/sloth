package api

import (
	"encoding/json"
	"github.com/ckeyer/sloth/types"
	"net/http"
)

func Login(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	u := new(types.User)
	err := json.NewDecoder(req.Body).Decode(u)
	if err != nil {
		ctx.Error()
	}
}

func Logout(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {

}

func Registry(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {

}
