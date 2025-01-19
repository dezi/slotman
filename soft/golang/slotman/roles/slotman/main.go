package main

import (
	"slotman/services/impl/provider"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/daemon"
	"slotman/utils/exitter"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	//_ = logger.StartService()
	_ = provider.StartService()

	gc9a01.TestDisplay()

	_ = exitter.StartService()

	_ = provider.StopService()
	//_ = logger.StopService()
}
