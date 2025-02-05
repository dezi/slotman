package tacho

import (
	"slotman/services/iface/tacho"
	"slotman/services/impl/provider"
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sync"
	"time"
)

type Service struct {
	tachoSensor *mcp23017.MCP23017
	tachoChan   chan SpeedRead
	tachoStates map[int]SpeedState

	trackStates map[int]TrackState

	mapsLock  sync.Mutex
	waitGroup sync.WaitGroup

	isProxyServer bool
	isProxyClient bool

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

	singleTon.tachoChan = make(chan SpeedRead, 10)
	singleTon.tachoStates = make(map[int]SpeedState)
	singleTon.trackStates = make(map[int]TrackState)

	singleTon.isProxyServer = simple.GetExecName() == "proxy"
	singleTon.isProxyClient = simple.GOOS == "darwin"

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

	if singleTon.tachoSensor != nil {
		_ = singleTon.tachoSensor.Close()
		singleTon.tachoSensor = nil
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
