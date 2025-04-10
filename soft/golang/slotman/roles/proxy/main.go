package main

import (
	"slotman/services/impl/provider"
	"slotman/services/impl/proxy"
	"slotman/services/impl/speedi"
	"slotman/services/impl/tacho"
	"slotman/services/impl/turner"
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
	_ = speedi.StartService()
	_ = tacho.StartService()
	_ = turner.StartService()

	_ = exitter.WaitUntilTermination()

	_ = turner.StopService()
	_ = tacho.StopService()
	_ = speedi.StopService()
	_ = proxy.StopService()
	_ = provider.StopService()
}
