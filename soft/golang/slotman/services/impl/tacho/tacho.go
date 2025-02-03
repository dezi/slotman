package tacho

import (
	"slotman/services/iface/tacho"
	"slotman/services/impl/provider"
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"sync"
	"time"
)

type Service struct {
	speedSensor *mcp23017.MCP23017
	speedChan   chan uint16
	speedStates map[int]SpeedState

	waitGroup sync.WaitGroup
	doExit    bool
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.speedChan = make(chan uint16, 10)
	singleTon.speedStates = make(map[int]SpeedState)

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
	singleTon.waitGroup.Wait()

	if singleTon.speedSensor != nil {
		_ = singleTon.speedSensor.Close()
		singleTon.speedSensor = nil
	}

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return tacho.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
