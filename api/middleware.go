package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/admin"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func CorsHandle(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Limt,Offset,Origin,Accept,X-Signature")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Max-Age", fmt.Sprint(24*time.Hour/time.Second))

	if ctx.Request.Method == "OPTIONS" {
		GinMessage(ctx, 200, "ok")
	}
}

func GinLogger(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()

	log.WithFields(log.Fields{
		"Method": ctx.Request.Method,
		"URL":    ctx.Request.URL.Path,
		"Status": ctx.Writer.Status(),
		"Period": fmt.Sprintf("%.6f", time.Now().Sub(start).Seconds()),
	}).Debug("bye jack.")

}

func MWNeedLogin(ctx *gin.Context) {
	xsign := ctx.Request.Header.Get("X-Signature")
	if xsign == "" {
		var err error
		xsign, err = ctx.Cookie("x-signature")
		if err != nil {
			GinError(ctx, 401, "Not Found Header X-Signature")
			return
		}
	}

	signSli := strings.Split(xsign, ":")
	if len(signSli) != 3 {
		GinError(ctx, 401, "Invalid Header X-Signature")
		return
	}

	apiKey, timestamp, sign := signSli[0], signSli[1], signSli[2]
	db := ctx.MustGet(CtxMgoDB).(*mgo.Database)

	ua, err := admin.AuthSignature(db, apiKey, timestamp, sign)
	if err != nil {
		GinError(ctx, 401, "Invalid signature content")
		return
	}

	ctx.Set(CtxUserAuth, ua)
}

func MWAuthGithubServer(rw http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("first read body error, ", err)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	key := []byte("asdf")
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	expectedMAC := mac.Sum(nil)
	if fmt.Sprintf("sha1=%x", expectedMAC) != req.Header.Get("X-Hub-Signature") {
		log.Warningf("the webhooks' sha1 from github should be %s, but now is %x",
			req.Header.Get("X-Hub-Signature"), expectedMAC)
	}
	log.Debugf("github server auth passing")
}
