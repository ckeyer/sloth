package main

import (
	logging "github.com/ckeyer/go-log"
)

//
func main2() {
	log := logging.GetDefaultLogger("ckeyer test")

	log.Debug("heelosdf")
	log.Error("heleo")
	log.Info("helelo")
	log.Warning("hhhhh")
}
