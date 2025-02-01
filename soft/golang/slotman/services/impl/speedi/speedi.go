package speedi

import (
	"slotman/services/iface/speedi"
	"slotman/services/impl/provider"
	"slotman/things/ti/ads1115"
	"slotman/utils/log"
	"time"
)

type Service struct {
	ads1115s []*ads1115.ADS1115

	ads1115Device1 *ads1115.ADS1115
	ads1115Device2 *ads1115.ADS1115

	speedControlAttached     []bool
	speedControlChannels     []chan uint16
	speedControlCalibrations []*SpeedControlCalibration

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

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return speedi.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
