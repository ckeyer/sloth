package api

import (
	"encoding/json"
	"fmt"

	"github.com/ckeyer/sloth/account"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func Login(ctx *gin.Context) {
	u := new(account.User)
	err := json.NewDecoder(ctx.Request.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)
	ret, err := u.Login(db)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	ss := sessions.Default(ctx)
	ss.Set("user", ret)

	ctx.JSON(200, ret)
}

func Logout(ctx *gin.Context) {
	GinMessage(ctx, 500, "...")
}

func Registry(ctx *gin.Context) {
	u := new(account.User)
	err := json.NewDecoder(ctx.Request.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	if u.Name == "" || u.Email == "" || len(u.Password) == 0 {
		GinError(ctx, 400, fmt.Errorf("Name, Email or Password not be nil"))
		return
	}

	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)
	ret, err := u.Registry(db)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	ctx.JSON(201, ret)
}
