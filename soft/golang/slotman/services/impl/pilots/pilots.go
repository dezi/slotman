package pilots

import (
	"image"
	"slotman/services/iface/pilots"
	"slotman/services/iface/teams"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sync"
	"time"
)

type Service struct {
	tms teams.Interface

	pilots map[simple.UUIDHex]*slotman.Pilot

	pilotProfileFull  map[simple.UUIDHex]*image.RGBA
	pilotProfileSmall map[simple.UUIDHex]*image.RGBA

	mapsLock sync.Mutex
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.tms, err = teams.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.pilots = make(map[simple.UUIDHex]*slotman.Pilot)
	singleTon.pilotProfileFull = make(map[simple.UUIDHex]*image.RGBA)
	singleTon.pilotProfileSmall = make(map[simple.UUIDHex]*image.RGBA)

	singleTon.loadMockups()

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return pilots.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
