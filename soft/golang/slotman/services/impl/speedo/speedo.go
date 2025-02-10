package speedo

import (
	"slotman/services/iface/race"
	"slotman/services/iface/speedo"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/things/pololu/mxt550"
	"slotman/utils/log"
	"time"
)

type Service struct {
	rce race.Interface

	tracksEnable []bool

	mxt550s []*mxt550.MXT550

	mxt550Motoron1 *mxt550.MXT550
	mxt550Motoron2 *mxt550.MXT550
	mxt550Motoron3 *mxt550.MXT550
	mxt550Motoron4 *mxt550.MXT550

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

	singleTon.rce, err = race.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.tracksEnable = make([]bool, slotman.MaxTracks)

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
	return speedo.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
