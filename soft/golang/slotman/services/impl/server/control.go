package server

import (
	"slotman/services/iface/pilots"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

func (sv *Service) DoControlTask() {
	sv.checkSetup()
	sv.checkServer()
}

func (sv *Service) checkSetup() {

	if sv.setup != nil {
		return
	}

	var plt pilots.Interface
	var tryErr error

	for {

		plt, tryErr = pilots.GetInstance()
		if tryErr == nil {
			break
		}

		time.Sleep(time.Millisecond * 250)
	}

	sv.setup = &slotman.Setup{

		Race: slotman.Race{
			What: "race",
			Mode: "set",

			Title:  "Großer Preis von Dezi",
			Tracks: 2,
			Rounds: 10,
		},

		Tracks: slotman.Tracks{
			What: "tracks",
			Mode: "set",

			Tracks: 2,
		},

		Pilots: make(map[simple.UUIDHex]*slotman.Pilot),
	}

	allPilots := plt.GetAllPilots()

	for _, pilot := range allPilots {
		sv.setup.Pilots[pilot.Uuid] = pilot
	}
}

func (sv *Service) checkServer() {
	err := sv.startServers()
	log.Cerror(err)
}
