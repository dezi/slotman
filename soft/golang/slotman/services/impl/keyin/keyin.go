package keyin

import (
	"os"
	keyinIface "slotman/services/iface/keyin"
	"slotman/services/impl/provider"
	keyinTypes "slotman/services/type/keyin"
	"slotman/utils/log"
	"sync"
	"time"
)

type Service struct {
	consoleReader *os.File

	subscribers map[keyinTypes.Subscriber]bool

	subscribersLock sync.Mutex

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

	singleTon.subscribers = make(map[keyinTypes.Subscriber]bool)

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
	return keyinIface.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
