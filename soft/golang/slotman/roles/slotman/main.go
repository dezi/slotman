package main

import (
	"slotman/utils/daemon"
	"slotman/utils/exitter"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	//_ = logger.StartService()

	_ = exitter.StartService()

	//_ = logger.StopService()
}
