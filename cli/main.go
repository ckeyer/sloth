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

const (
	appName = "sloth"
)

var (
	// for flags.
	debug                     bool
	addr, raddr, rauth, uiDir string // runCmd
	outputFile                string // evalCmd

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

	evaluateCmd = &cli.Command{
		Name:    "eval",
		Aliases: []string{"evaluate", "e"},
		Usage:   "Evaluate one golang file or project, and generate a html file.",
		Flags: []cli.Flag{
			debugFlag,
			outputFlag,
		},
		Action: func(ctx *cli.Context) error {

			return nil
		},
	}
)

func main() {
	app := &cli.App{
		Name:    appName,
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
				log.Debugf("%s at debug model.", appName)
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
