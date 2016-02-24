package app

import (
	"fmt"
	"strings"
)

type App struct {
	Command string
	Args    []string
}

func init() {
	config = newFlag("config", "c", "config.json")
}

func NewApp(args []string) (app *App, err error) {
	if len(args) <= 0 {
		return
	}
	app = new(App)

	err = app.paramArgs(args)
	return app, err
}

func (a *App) paramArgs(args []string) error {
	for i := 0; i < len(args); i += 2 {
		switch {
		case strings.HasPrefix(args[i], "--"):
			if len(args[i]) > 2 && len(args[i:]) >= 2 {
				a.Config[args[i][2:]] = args[i+1]
			} else {
				return fmt.Errorf("%s error", args[i:])
			}
		case strings.HasPrefix(args[i], "-"):
			if len(args[i]) > 1 && len(args[i:]) >= 2 {
				a.Config[args[i][1:]] = args[i+1]
			} else {
				return fmt.Errorf("%s error", args[i:])
			}
		default:
			if len(args[i:]) > 1 {
				a.Command = args[i]
				a.Args = args[i+1:]
			} else if len(args[i:]) == 1 {
				a.Command = args[i]
			} else {
				return fmt.Errorf("length error ")
			}
		}
	}
	return nil
}

func (a *App) paramFlags() {

}
