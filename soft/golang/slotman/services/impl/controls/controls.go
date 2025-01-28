package controls

import (
	"slotman/services/iface/controls"
	"slotman/services/impl/provider"
	"slotman/things/mcp/mcp23017"
	"slotman/things/pololu/mxt550"
	"slotman/things/ti/ads1115"
	"slotman/utils/log"
	"time"
)

type Service struct {
	ads1115s  []*ads1115.ADS1115
	mxt550s   []*mxt550.MXT550
	mcp23017s []*mcp23017.MCP23017

	ads1115Device1 *ads1115.ADS1115
	ads1115Device2 *ads1115.ADS1115

	mxt550Motoron1 *mxt550.MXT550
	mxt550Motoron2 *mxt550.MXT550
	mxt550Motoron3 *mxt550.MXT550

	mcp23017StartLight   *mcp23017.MCP23017
	mcp23017SpeedMeasure *mcp23017.MCP23017

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
	return controls.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
