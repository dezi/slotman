package main

import (
	"slotman/services/impl/provider"
	"slotman/services/impl/slotdisplay"
	"slotman/utils/daemon"
	"slotman/utils/exitter"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	//_ = logger.StartService()
	_ = provider.StartService()

	_ = slotdisplay.StartService()

	_ = exitter.StartService()

	_ = slotdisplay.StopService()

	_ = provider.StopService()
	//_ = logger.StopService()
}
