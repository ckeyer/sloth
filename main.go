package main

import (
	"github.com/ckeyer/go-ci/api"
	"github.com/ckeyer/go-ci/lib"
	_ "github.com/ckeyer/go-ci/types"
	"gopkg.in/mgo.v2/bson"
)

var (
	log = lib.GetLogger()

	host     = "192.168.2.230"
	port     = "27017"
	database = "go-ci"
	userName = ""
	password = ""
)

func init() {
	db := lib.NewMdb(host, port, database, userName, password)
	db.Insert("gogogo", bson.M{"name": "wang"})
}

func main() {
	log.Notice("server is running...")
	api.Serve(":8080")
}
