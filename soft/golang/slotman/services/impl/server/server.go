package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"slotman/services/iface/server"
	"slotman/services/impl/provider"
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

	webSockets map[simple.UUIDHex]*websocket.Conn
	mapsLock   sync.Mutex
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.webSockets = make(map[simple.UUIDHex]*websocket.Conn)

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
