package server

import (
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) DoControlTask() {
	sv.checkSetup()
	sv.checkServer()
}

func (sv *Service) checkSetup() {

	if sv.setup != nil {
		return
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
}

func (sv *Service) checkServer() {
	err := sv.startServers()
	log.Cerror(err)
}
