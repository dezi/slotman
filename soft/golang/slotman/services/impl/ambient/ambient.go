package ambient

import (
	"slotman/services/iface/ambient"
	"slotman/services/impl/provider"
	"slotman/things/sensirion/sgp40"
	"slotman/utils/log"
	"time"
)

type Service struct {
	sgp40Co2 *sgp40.SGP40

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

	if singleTon.sgp40Co2 != nil {
		_ = singleTon.sgp40Co2.Close()
		singleTon.sgp40Co2 = nil
	}

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return ambient.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
