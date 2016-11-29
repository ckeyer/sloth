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
	err := json.NewDecoder(ctx.Request.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	_ = u
}

func Logout(rw http.ResponseWriter, req *http.Request) {

}

func Registry(ctx *gin.Context) {
	u := new(admin.User)
	err := json.NewDecoder(ctx.Request.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	if u.Email == "" || len(u.Password) == 0 {
		GinError(ctx, 400, fmt.Errorf("Email or Password not be nil"))
		return
	}
	GinMessage(ctx, 201, u)
}
