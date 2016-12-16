package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func NewGin() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())

	return g
}

// transform (http).Handler to Gin HandleFunc
func GinH(h interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ghf, ok := h.(gin.HandlerFunc); ok {
			ghf(ctx)
			return
		}
		if h, ok := h.(http.HandlerFunc); ok {
			h(ctx.Writer, ctx.Request)
			return
		}
		if ser, ok := h.(http.Handler); ok {
			ser.ServeHTTP(ctx.Writer, ctx.Request)
			return
		}
		log.Errorf("cannot transform %T --> gin.HandlerFunc", h)
	}
}

func SetMgoDB(db *mgo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newDb := db.Session.Clone().DB(db.Name)
		ctx.Set(CtxMgoDB, newDb)
		ctx.Next()
		newDb.Session.Close()
		log.Debugf("closed mongo session.")
	}
}

func GinError(ctx *gin.Context, status int, err ...interface{}) {
	defer func() {
		ctx.Abort()
	}()

	l := len(err)
	ret := map[string]interface{}{"error": ""}
	if l >= 1 {
		ret["error"] = fmt.Sprint(err[0])
	}
	if l >= 2 {
		ret["message"] = err[1:]
	}

	ctx.JSON(status, ret)
}

func GinMessage(ctx *gin.Context, status int, msg ...interface{}) {
	defer func() {
		ctx.Abort()
	}()

	l := len(msg)
	ret := map[string]interface{}{"message": ""}
	if l >= 1 {
		ret["message"] = fmt.Sprint(msg[0])
		log.WithField("message", msg[0]).Debug("return message.")
	}
	if l >= 2 {
		ret["supplement"] = msg[1:]
	}

	ctx.JSON(status, ret)
}
