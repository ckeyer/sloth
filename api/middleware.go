package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func CorsHandle(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Limt,Offset,Origin,Accept")
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
	token := ctx.Request.Header.Get("X-Token")
	if token == "" {
		var err error
		token, err = ctx.Cookie("UserToken")
		if err != nil {
			GinError(ctx, 401, "Not Found Header X-Token")
			return
		}
	}
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
