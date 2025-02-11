package race

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/pilots"
	"slotman/services/iface/race"
	"slotman/services/iface/server"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/iface/tacho"
	"slotman/services/iface/teams"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"time"
)

type Service struct {
	srv server.Interface
	amp ampel.Interface
	sdi speedi.Interface
	sdo speedo.Interface
	tco tacho.Interface
	tms teams.Interface
	plt pilots.Interface

	raceState     slotman.RaceState
	raceStateDone slotman.RaceState

	raceInfos []*slotman.RaceInfo

	tracksReady   []int
	tracksVoltage []int

	roundsToGo int

	servicesReady bool
	looperStarted bool

	doExit bool
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.raceState = slotman.RaceStateIdle

	singleTon.tracksReady = make([]int, slotman.MaxTracks)
	singleTon.tracksVoltage = make([]int, slotman.MaxTracks)

	singleTon.raceInfos = make([]*slotman.RaceInfo, slotman.MaxTracks)

	for track := range singleTon.raceInfos {
		singleTon.raceInfos[track] = &slotman.RaceInfo{
			What:  "info",
			Mode:  "set",
			Track: track,
		}
	}

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	singleTon.doExit = true

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return race.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
