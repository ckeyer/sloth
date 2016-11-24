package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/api"
	_ "github.com/ckeyer/sloth/types"
	"github.com/ckeyer/sloth/version"
	_ "github.com/fsouza/go-dockerclient"
	_ "gopkg.in/check.v1"
	"gopkg.in/urfave/cli.v2"
)

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
