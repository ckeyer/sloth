package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func MWHello(rw http.ResponseWriter, req *http.Request) {
	log.Debugf("hello %s:%s", req.Method, req.URL.Path)
}

func WMAuthGithubServer(rw http.ResponseWriter, req *http.Request) {
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
