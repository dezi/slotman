package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	serverIface "slotman/services/iface/server"
	"slotman/services/impl/provider"
	serverTypes "slotman/services/type/server"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sync"
	"time"
)

type Service struct {
	httpMux     *http.ServeMux
	httpServer  *http.Server
	httpRunning bool

	webClientsConns      map[simple.UUIDHex]*websocket.Conn
	webClientsConnsLocks map[simple.UUIDHex]*sync.Mutex
	webClientsLock       sync.Mutex

	subscribers     map[string]serverTypes.Subscriber
	subscribersLock sync.Mutex

	mapsLock sync.Mutex

	setup *slotman.Setup
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.webClientsConns = make(map[simple.UUIDHex]*websocket.Conn)
	singleTon.webClientsConnsLocks = make(map[simple.UUIDHex]*sync.Mutex)

	singleTon.subscribers = make(map[string]serverTypes.Subscriber)

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
	return serverIface.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
