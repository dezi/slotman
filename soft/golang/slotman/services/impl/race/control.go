package race

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/pilots"
	"slotman/services/iface/server"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/iface/tacho"
	"slotman/services/iface/teams"
	"slotman/services/type/slotman"
	"slotman/utils/simple"
	"time"
)

func (sv *Service) DoControlTask() {
	sv.checkServices()
	sv.checkSetup()
	sv.checkLooper()

	//log.Printf("Race raceState=%s", sv.raceState)
}

func (sv *Service) checkServices() {

	if sv.servicesReady {
		return
	}

	//
	// Yield time for target services to register.
	//

	time.Sleep(time.Second)

	var tryErr error

	sv.srv, tryErr = server.GetInstance()
	if tryErr != nil {
		return
	}

	sv.amp, tryErr = ampel.GetInstance()
	if tryErr != nil {
		return
	}

	sv.sdi, tryErr = speedi.GetInstance()
	if tryErr != nil {
		return
	}

	sv.sdo, tryErr = speedo.GetInstance()
	if tryErr != nil {
		return
	}

	sv.tco, tryErr = tacho.GetInstance()
	if tryErr != nil {
		return
	}

	sv.plt, tryErr = pilots.GetInstance()
	if tryErr != nil {
		return
	}

	sv.tms, tryErr = teams.GetInstance()
	if tryErr != nil {
		return
	}

	sv.servicesReady = true
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

func (sv *Service) checkLooper() {

	if !sv.servicesReady || sv.looperStarted {
		return
	}

	sv.looperStarted = true

	go sv.looper()
}
