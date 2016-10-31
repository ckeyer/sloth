package main

import (
	"github.com/ckeyer/sloth/api"
	_ "github.com/ckeyer/sloth/types"
	"github.com/ckeyer/sloth/utils"
	"gopkg.in/mgo.v2/bson"

	_ "github.com/fsouza/go-dockerclient"
	_ "gopkg.in/check.v1"
)

var (
	log = lib.GetLogger()

	host     = "192.168.2.230"
	port     = "27017"
	database = "go-ci"
	userName = ""
	password = ""
)

func init2() {
	// mongodb initialization
	db := lib.NewMdb(host, port, database, userName, password)
	db.Insert("gogogo", bson.M{"name": "wang"})

	// redis initialization
	lib.InitRedis("127.0.0.1", "6379")
}

func main() {
	// api.SetStatus()
	// return
	log.Notice("server is running...")
	api.Serve(":8080")
}
