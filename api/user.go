package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ckeyer/sloth/admin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func Login(ctx *gin.Context) {
	u := new(admin.User)
	err := json.NewDecoder(ctx.Request.Body).Decode(u)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)
	user, err := u.Login(db)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	ss := sessions.Default(ctx)
	ss.Set("user", user)

	ua := admin.NewUserAuth(user.ID, time.Now().Add(time.Hour*24*90))
	if err := ua.Insert(db); err != nil {
		GinError(ctx, 500, err)
		return
	}

	ret := map[string]interface{}{
		"user":      user,
		"user_auth": ua,
	}
	ctx.JSON(200, ret)
}

func Logout(ctx *gin.Context) {
	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)
	ua := ctx.MustGet(CtxUserAuth).(*admin.UserAuth)

	if err := ua.Remove(db); err != nil {
		GinError(ctx, 500, err)
		return
	}

	GinMessage(ctx, 200, "ok")
}

func Registry(ctx *gin.Context) {
	u := new(admin.User)
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
