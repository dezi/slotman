package race

import (
	"slotman/services/iface/ampel"
	raceIface "slotman/services/iface/race"
	"slotman/services/iface/speedo"
	"slotman/services/impl/provider"
	raceTypes "slotman/services/type/race"
	"slotman/utils/log"
	"time"
)

type Service struct {
	amp ampel.Interface
	sdo speedo.Interface

	raceState     raceTypes.RaceState
	raceStateDone raceTypes.RaceState

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
	interval = time.Second * 60
	return
}
