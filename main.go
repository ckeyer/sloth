package main

import (
	"github.com/ckeyer/go-ci/api"
	"github.com/ckeyer/go-ci/lib"
	_ "github.com/ckeyer/go-ci/types"
)

var log = lib.GetLogger()

func init() {

}

func main() {
	log.Notice("server is running...")
	api.Serve(":8080")
}
