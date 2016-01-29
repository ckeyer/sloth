package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ckeyer/go-sh/buffer"
)

func main() {
	cmd := exec.Command("sh", "test.sh")

	buf := buffer.NewBuf()
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
