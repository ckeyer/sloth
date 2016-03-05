package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/go-ci/api"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	api.SetStatus()
	return
	api.Serve(":8080")
}
