package ampel

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/race"
	"slotman/services/impl/provider"
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"sync"
	"time"
)

type Service struct {
	rce race.Interface

	ampelGpio *mcp23017.MCP23017
	ampelLock sync.Mutex

	waitingTracks      int
	waitingTracksReady []int

	ampelState AmpelState
	roundsToGo int

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

	if singleTon.ampelGpio != nil {
		_ = singleTon.ampelGpio.Close()
		singleTon.ampelGpio = nil
	}

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return ampel.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
