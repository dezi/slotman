package pilots

import (
	"image"
	"slotman/services/iface/pilots"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sync"
	"time"
)

type Service struct {
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

	singleTon.pilots = make(map[simple.UUIDHex]*slotman.Pilot)
	singleTon.pilotProfileFull = make(map[simple.UUIDHex]*image.RGBA)
	singleTon.pilotProfileSmall = make(map[simple.UUIDHex]*image.RGBA)

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

func (sv *Service) GetName() (name provider.Provider) {
	return pilots.Provider
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
