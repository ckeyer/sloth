package api

import (
	"encoding/json"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/admin"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// GET /settings/:key
func GetSettings(ctx *gin.Context) {
	var keys []string
	if ctx.Param("key") != "" {
		keys = []string{ctx.Param("key")}
	} else {
		keys = strings.Split(ctx.Query("key"), ",")
	}
	if len(keys) < 1 || keys[0] == "" {
		GinError(ctx, 400, "invalid key")
		return
	}
	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)

	log.Debugf("get settings, %+v", keys)
	ret, err := admin.GetValues(db, keys...)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, ret)
}

// POST /settings
// body: ["key": "value"]
func AddSettings(ctx *gin.Context) {
	kv := map[string]string{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&kv)
	if err != nil {
		log.Errorf("invalid body, %s", err)
		GinError(ctx, 400, "invalid body")
		return
	}

	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	for k, v := range kv {
		err = admin.SetKV(db, k, v)
		if err != nil {
			GinError(ctx, 500, err)
			return
		}
	}

	ctx.Status(204)
}
