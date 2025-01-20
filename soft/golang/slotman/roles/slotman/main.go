package main

import (
	"slotman/services/impl/pilots"
	"slotman/services/impl/provider"
	"slotman/services/impl/turner"
	"slotman/utils/daemon"
	"slotman/utils/exitter"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	//_ = logger.StartService()
	_ = provider.StartService()

	_ = pilots.StartService()
	_ = turner.StartService()

	_ = exitter.StartService()

	_ = pilots.StopService()
	_ = turner.StopService()

	_ = provider.StopService()
	//_ = logger.StopService()
}
