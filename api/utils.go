package api

import (
	"io/ioutil"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/ckeyer/sloth/pkgs/admin"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func ping(ctx *gin.Context) {
	GinMessage(ctx, 200, "hi")
}

// GetStatus GET /status
func GetStatus(ctx *gin.Context) {
	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	ret, err := admin.Status(db)
	if err != nil {
		log.Error("get status failed", err)
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, ret)
}

// TODO ...
func TODO(ctx *gin.Context) {
	bs, _ := ioutil.ReadAll(ctx.Request.Body)

	log.WithFields(log.Fields{
		"Method":  ctx.Request.Method,
		"Path":    ctx.Request.URL.Path,
		"Remote":  ctx.Request.RemoteAddr,
		"Headers": ctx.Request.Header,
		"Agent":   ctx.Request.UserAgent(),
	}).Debug(string(bs))
	GinMessage(
		ctx,
		503,
		"Function is under development...",
		ctx.Request.Method,
		ctx.Request.URL.Path,
	)
}

func getListQuery(ctx *gin.Context) (offset, limit int) {
	offstr := ctx.Query("offset")
	limstr := ctx.Query("limit")
	if offstr == "" {
		offstr = ctx.Request.Header.Get("Offset")
	}
	if limstr == "" {
		limstr = ctx.Request.Header.Get("Limit")
	}

	offset, _ = strconv.Atoi(offstr)
	limit, _ = strconv.Atoi(limstr)

	if offset < 0 {
		offset = 0
	}
	if limit <= 0 && limit > 50 {
		limit = 50
	}

	return offset, limit
}
