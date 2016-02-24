package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	libapp "github.com/ckeyer/go-ci/app"
	"github.com/ckeyer/go-ci/types"
)

var (
	app *libapp.App
)

func init() {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatalln("no args...")
	}
	var err error
	app, err = libapp.NewApp(args)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(app)
}

func main() {

	return
	cmd := exec.Command("sh", "test.sh")

	buf := types.NewBuf()
	buf.SyncRead = true

	cmd.Stderr = os.Stderr
	cmd.Stdout = buf
	cmd.Start()

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait... err, ", err)
	} else {
		fmt.Println("wait...")
		buf.Over()
		fmt.Println("input times, ", buf.GetInputTimes())
	}
}
