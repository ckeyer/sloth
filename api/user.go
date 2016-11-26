package api

import (
	"encoding/json"
	"net/http"

	"github.com/ckeyer/sloth/types"
)

func Login(rw http.ResponseWriter, req *http.Request) {
	u := new(types.User)
	json.NewDecoder(req.Body).Decode(u)
	// if err != nil {
	// 	ctx.Error()
	// }
}

func Logout(rw http.ResponseWriter, req *http.Request) {

}

func Registry(rw http.ResponseWriter, req *http.Request) {

}
