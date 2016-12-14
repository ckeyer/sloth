package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/api"
	"github.com/ckeyer/sloth/global"
	"github.com/ckeyer/sloth/version"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/check.v1"
	"gopkg.in/urfave/cli.v2"
)

var (
	// for flags.
	debug                             bool
	addr, raddr, rauth, uiDir, mgoURL string // runCmd
	outputFile                        string // evalCmd

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
		},
		Before: func(ctx *cli.Context) error {
			if mgoURL == "" {
				return fmt.Errorf("invalid flags mgo_url(ENV MGO_URL)")
			}

			return nil
		},
		Action: func(ctx *cli.Context) error {
			db, err := global.InitDB(mgoURL)
			if err != nil {
				return err
			}
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
		},
		Action: func(ctx *cli.Context) error {
			return cli.ShowAppHelp(ctx)
		},
	}

	app.Run(os.Args)
}
