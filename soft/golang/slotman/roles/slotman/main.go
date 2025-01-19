package main

import (
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/daemon"
	"slotman/utils/exitter"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	//_ = logger.StartService()

	gc9a01.TestDisplay()

	_ = exitter.StartService()

	//_ = logger.StopService()
}
