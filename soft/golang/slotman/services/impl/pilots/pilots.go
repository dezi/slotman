package pilots

import (
	"slotman/services/iface/slotdisplay"
	"slotman/services/impl/provider"
	"slotman/utils/log"
	"time"
)

type Service struct {
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

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
	return slotdisplay.Provider
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
