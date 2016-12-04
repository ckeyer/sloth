package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/api"
	_ "github.com/ckeyer/sloth/types"
	"github.com/ckeyer/sloth/version"
	_ "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/check.v1"
	"gopkg.in/urfave/cli.v2"
)

var (
	// for flags.
	debug                     bool
	addr, raddr, rauth, uiDir string // runCmd

	// Flags
	addrFlag = &cli.StringFlag{
		Name:        "addr",
		Aliases:     []string{"address"},
		EnvVars:     []string{"ADDR"},
		Value:       ":8000",
		Destination: &addr,
	}
	raddrFlag = &cli.StringFlag{
		Name:        "redis_addr",
		Aliases:     []string{"raddr"},
		EnvVars:     []string{"REDIS_ADDR"},
		Value:       "127.0.0.1:6379",
		Destination: &raddr,
	}
	rauthFlag = &cli.StringFlag{
		Name:        "redis_auth",
		Aliases:     []string{"rauth"},
		EnvVars:     []string{"REDIS_AUTH"},
		Value:       "",
		Destination: &rauth,
	}
	uiDirFlag = &cli.StringFlag{
		Name:        "ui_dir",
		Aliases:     []string{"uiDir"},
		EnvVars:     []string{"UI_DIR"},
		Value:       "./assets",
		Destination: &uiDir,
	}
	debugFlag = &cli.BoolFlag{
		Name:        "debug",
		Aliases:     []string{"D"},
		EnvVars:     []string{"DEBUG"},
		Value:       false,
		Destination: &debug,
	}

	// Authors
	ckeyer = &cli.Author{
		Name:  "ckeyer",
		Email: "me@ckeyer.com",
	}

	// Commands
	runCmd = &cli.Command{
		Name: "run",
		Flags: []cli.Flag{
			addrFlag,
			raddrFlag,
			rauthFlag,
			uiDirFlag,
			debugFlag,
		},
		Action: func(ctx *cli.Context) error {
			log.Info("server is running at ", addr)
			api.Serve(addr)
			return nil
		},
	}
)

func main() {

	app := &cli.App{
		Name:    "sloth",
		Version: version.GetVersion(),
		Usage:   "",
		Authors: []*cli.Author{
			ckeyer,
		},
		Flags: []cli.Flag{
			debugFlag,
		},
		Before: func(ctx *cli.Context) error {
			/// init config model.
			gin.SetMode(gin.ReleaseMode)
			log.SetFormatter(&log.JSONFormatter{})
			if debug {
				log.SetLevel(log.DebugLevel)
				log.Debug("server is running at debug model.")
			} else {
				log.SetLevel(log.InfoLevel)
			}

			return nil
		},
		Commands: []*cli.Command{
			runCmd,
		},
		Action: func(ctx *cli.Context) error {
			return cli.ShowAppHelp(ctx)
		},
	}

	app.Run(os.Args)
}
