package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ckeyer/sloth/admin"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	u := new(admin.User)
	err := json.NewDecoder(req.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	_ = u
}

func Logout(rw http.ResponseWriter, req *http.Request) {

}

func Registry(rw http.ResponseWriter, req *http.Request) error {
	u := new(admin.User)
	err := json.NewDecoder(req.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	if u.Email == "" || u.Password == "" {
		return fmt.Errorf("Email or Password not be nil")
	}
}
