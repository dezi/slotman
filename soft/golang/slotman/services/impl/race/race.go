package race

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/pilots"
	raceIface "slotman/services/iface/race"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/iface/teams"
	"slotman/services/impl/provider"
	raceTypes "slotman/services/type/race"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"time"
)

type Service struct {
	amp ampel.Interface
	sdi speedi.Interface
	sdo speedo.Interface
	tms teams.Interface
	plt pilots.Interface

	raceState     raceTypes.RaceState
	raceStateDone raceTypes.RaceState
	raceRecords   []raceTypes.RaceRecord

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

	singleTon.raceState = raceTypes.RaceStateIdle

	singleTon.tracksReady = make([]int, slotman.MaxTracks)
	singleTon.tracksVoltage = make([]int, slotman.MaxTracks)
	singleTon.raceRecords = make([]raceTypes.RaceRecord, slotman.MaxTracks)

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
	return raceIface.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
