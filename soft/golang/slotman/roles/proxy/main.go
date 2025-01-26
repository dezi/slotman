package main

import (
	"slotman/services/impl/provider"
	"slotman/services/impl/proxy"
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

	_ = exitter.WaitUntilTermination()

	_ = proxy.StopService()
	_ = provider.StopService()
}
