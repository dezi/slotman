package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"slotman/services/iface/server"
	"slotman/services/impl/provider"
	proxyTypes "slotman/services/type/proxy"
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

	setup *slotman.Setup

	webClientsConns      map[simple.UUIDHex]*websocket.Conn
	webClientsConnsLocks map[simple.UUIDHex]*sync.Mutex
	webClientsLock       sync.Mutex

	subscribers     map[string]proxyTypes.Subscriber
	subscribersLock sync.Mutex

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

	singleTon.webClientsConns = make(map[simple.UUIDHex]*websocket.Conn)
	singleTon.webClientsConnsLocks = make(map[simple.UUIDHex]*sync.Mutex)

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
	return server.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
