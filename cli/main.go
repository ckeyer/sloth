package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/api"
	_ "github.com/ckeyer/sloth/types"
	"github.com/ckeyer/sloth/utils"
	_ "github.com/fsouza/go-dockerclient"
	_ "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

var (
	host     = "192.168.2.230"
	port     = "27017"
	database = "go-ci"
	userName = ""
	password = ""
)

func init2() {
	// mongodb initialization
	db := utils.NewMdb(host, port, database, userName, password)
	db.Insert("gogogo", bson.M{"name": "wang"})

	// redis initialization
	utils.InitRedis("127.0.0.1", "6379")
}

func main() {
	log.Info("server is running...")
	api.Serve(":8080")
}
