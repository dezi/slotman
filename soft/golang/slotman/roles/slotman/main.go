package main

import (
	"slotman/services/impl/provider"
	"slotman/services/impl/proxy"
	"slotman/services/impl/tacho"
	"slotman/utils/daemon"
	"slotman/utils/exitter"
	"slotman/utils/log"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	log.SetCallerLength(48)

	_ = provider.StartService()
	_ = proxy.StartService()

	//_ = teams.StartService()
	//_ = pilots.StartService()
	//_ = turner.StartService()
	//_ = controls.StartService()
	//_ = speedo.StartService()
	//_ = speedi.StartService()
	_ = tacho.StartService()

	_ = exitter.WaitUntilTermination()

	_ = tacho.StopService()
	//_ = speedi.StopService()
	//_ = speedo.StopService()
	//_ = controls.StopService()
	//_ = turner.StopService()
	//_ = pilots.StopService()
	//_ = teams.StopService()

	_ = proxy.StopService()
	_ = provider.StopService()
}
