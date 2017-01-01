package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/api"
	"github.com/ckeyer/sloth/docker"
	"github.com/ckeyer/sloth/global"
	"github.com/ckeyer/sloth/version"
	"github.com/ckeyer/sloth/views"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/check.v1"
	"gopkg.in/urfave/cli.v2"
)

var (
	// for flags.
	debug bool
	addr,
	raddr,
	rauth,
	uiDir,
	mgoURL,
	dockerEndpoint string // runCmd
	outputFile string // evalCmd

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
	outputFlag = &cli.StringFlag{
		Name:        "output",
		Aliases:     []string{"o", "out"},
		EnvVars:     []string{"OUTPUT_FILE"},
		Value:       "sloth.html",
		Destination: &outputFile,
	}
	mgoURLFlag = &cli.StringFlag{
		Name:        "mgo_url",
		Aliases:     []string{"mongoUrl"},
		EnvVars:     []string{"MGO_URL"},
		Value:       "mongodb://localhost:27017/sloth",
		Destination: &mgoURL,
	}
	dockerEPFlag = &cli.StringFlag{
		Name:        "docker_ep",
		Aliases:     []string{"docker_endpoint"},
		EnvVars:     []string{"DOCKER_EP", "DOCKER_ENDPOINT"},
		Value:       "unix:///var/run/docker.sock",
		Destination: &dockerEndpoint,
	}

	// Authors
	ckeyer = &cli.Author{
		Name:  "ckeyer",
		Email: "me@ckeyer.com",
	}

	// Commands
	runCmd = &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Start web server.",
		Flags: []cli.Flag{
			addrFlag,
			uiDirFlag,
			debugFlag,
			mgoURLFlag,
			dockerEPFlag,
		},
		Before: func(ctx *cli.Context) error {
			if uiDir != "" {
				views.SetUIDir(uiDir)
			}

			if mgoURL == "" {
				return fmt.Errorf("invalid flags mgo_url(ENV MGO_URL)")
			}

			_, err := docker.Connect(dockerEndpoint)
			if err != nil {
				return err
			}

			return nil
		},
		Action: func(ctx *cli.Context) error {
			db, err := global.InitDB(mgoURL)
			if err != nil {
				return err
			}
			log.Info("connected mongodb successful.")
			log.Info("server is running at ", addr)
			api.Serve(addr, db)
			return nil
		},
	}

	evaluateCmd = &cli.Command{
		Name:    "eval",
		Aliases: []string{"evaluate", "e", "check"},
		Usage:   "Evaluate one golang file or project, and generate a html file.",
		Flags: []cli.Flag{
			debugFlag,
			outputFlag,
		},
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Action: func(ctx *cli.Context) error {
			log.Infof("args: %v", ctx.Args())
			return nil
		},
	}
)

func main() {

	app := &cli.App{
		Name:    "sloth",
		Version: version.GetCompleteVersion(),
		Usage:   "",
		Authors: []*cli.Author{
			ckeyer,
		},
		Flags: []cli.Flag{
			debugFlag,
		},
		Before: func(ctx *cli.Context) error {
			/// init config model.
			log.SetFormatter(&log.JSONFormatter{})
			if debug {
				gin.SetMode(gin.DebugMode)
				log.SetLevel(log.DebugLevel)
				log.Debug("server is running at debug model.")
			} else {
				gin.SetMode(gin.ReleaseMode)
				log.SetLevel(log.InfoLevel)
			}

			return nil
		},
		Commands: []*cli.Command{
			runCmd,
			evaluateCmd,
		},
		Action: func(ctx *cli.Context) error {
			return cli.ShowAppHelp(ctx)
		},
	}

	app.Run(os.Args)
}
