package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/api"
	_ "github.com/ckeyer/sloth/types"
	"github.com/ckeyer/sloth/utils"
	"github.com/ckeyer/sloth/version"
	_ "github.com/fsouza/go-dockerclient"
	_ "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/urfave/cli.v2"
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
	var debug bool
	var addr, raddr, rauth, uiDir string

	app := &cli.App{
		Name:    "sloth",
		Version: version.GetVersion(),
		Usage:   "",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "ckeyer",
				Email: "me@ckeyer.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				EnvVars:     []string{"DEBUG"},
				Value:       false,
				Destination: &debug,
			},
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"address"},
				EnvVars:     []string{"ADDR"},
				Value:       ":8080",
				Destination: &addr,
			},
			&cli.StringFlag{
				Name:        "redis_addr",
				Aliases:     []string{"raddr"},
				EnvVars:     []string{"REDIS_ADDR"},
				Value:       "127.0.0.1:6379",
				Destination: &raddr,
			},
			&cli.StringFlag{
				Name:        "redis_auth",
				Aliases:     []string{"rauth"},
				EnvVars:     []string{"REDIS_AUTH"},
				Value:       "",
				Destination: &rauth,
			},
			&cli.StringFlag{
				Name:        "ui_dir",
				Aliases:     []string{"uiDir"},
				EnvVars:     []string{"UI_DIR"},
				Value:       "./assets",
				Destination: &uiDir,
			},
		},
		Before: func(ctx *cli.Context) error {
			/// init config logger.
			log.SetFormatter(&log.JSONFormatter{})
			if debug {
				log.SetLevel(log.DebugLevel)
			} else {
				log.SetLevel(log.InfoLevel)
			}

			return nil
		},
		Action: func(ctx *cli.Context) error {
			log.Info("server is running at ", addr)
			api.Serve(addr)

			return nil
		},
	}

	app.Run(os.Args)
}
