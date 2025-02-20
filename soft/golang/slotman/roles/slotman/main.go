package main

import (
	"slotman/services/impl/ambient"
	"slotman/services/impl/ampel"
	"slotman/services/impl/identity"
	"slotman/services/impl/keyin"
	"slotman/services/impl/pilots"
	"slotman/services/impl/provider"
	"slotman/services/impl/proxy"
	"slotman/services/impl/race"
	"slotman/services/impl/server"
	"slotman/services/impl/speedi"
	"slotman/services/impl/speedo"
	"slotman/services/impl/storage"
	"slotman/services/impl/tacho"
	"slotman/services/impl/teams"
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
	_ = identity.StartService()
	_ = storage.StartService()
	_ = server.StartService()
	_ = proxy.StartService()
	_ = keyin.StartService()

	_ = race.StartService()
	_ = teams.StartService()
	_ = pilots.StartService()
	_ = speedo.StartService()
	_ = speedi.StartService()
	_ = tacho.StartService()
	_ = turner.StartService()
	_ = ampel.StartService()
	_ = ambient.StartService()

	_ = exitter.WaitUntilTermination()

	_ = ambient.StopService()
	_ = ampel.StopService()
	_ = turner.StopService()
	_ = tacho.StopService()
	_ = speedi.StopService()
	_ = speedo.StopService()
	_ = pilots.StopService()
	_ = teams.StopService()
	_ = race.StopService()

	_ = keyin.StopService()
	_ = proxy.StopService()
	_ = server.StopService()
	_ = storage.StopService()
	_ = identity.StopService()
	_ = provider.StopService()
}
