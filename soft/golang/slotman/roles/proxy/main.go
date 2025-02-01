package main

import (
	"slotman/services/impl/provider"
	"slotman/services/impl/proxy"
	"slotman/services/impl/speedi"
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

	_ = exitter.WaitUntilTermination()

	_ = speedi.StartService()
	_ = proxy.StopService()
	_ = provider.StopService()
}
