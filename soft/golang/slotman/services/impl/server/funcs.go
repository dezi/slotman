package server

import (
	"github.com/gorilla/websocket"
	"slotman/services/type/server"
	"slotman/utils/simple"
)

func (sv *Service) Subscribe(what string, handler server.Subscriber) {

	sv.subscribersLock.Lock()
	defer sv.subscribersLock.Unlock()

	sv.subscribers[what] = handler
}

func (sv *Service) Unsubscribe(what string) {

	sv.subscribersLock.Lock()
	defer sv.subscribersLock.Unlock()

	delete(sv.subscribers, what)
}

func (sv *Service) Transmit(appId simple.UUIDHex, resBytes []byte) (err error) {

	sv.webClientsLock.Lock()
	defer sv.webClientsLock.Unlock()

	webClientsConn := sv.webClientsConns[appId]
	webClientsConnLock := sv.webClientsConnsLocks[appId]

	if webClientsConn == nil || webClientsConnLock == nil {
		return
	}

	webClientsConnLock.Lock()
	err = webClientsConn.WriteMessage(websocket.TextMessage, resBytes)
	webClientsConnLock.Unlock()

	return
}

func (sv *Service) Broadcast(appId simple.UUIDHex, resBytes []byte) (err error) {

	sv.webClientsLock.Lock()
	defer sv.webClientsLock.Unlock()

	for appIdBc, webClientsConn := range sv.webClientsConns {

		if appId == appIdBc {
			continue
		}

		webClientsConnLock := sv.webClientsConnsLocks[appIdBc]

		if webClientsConn == nil || webClientsConnLock == nil {
			continue
		}

		webClientsConnLock.Lock()
		_ = webClientsConn.WriteMessage(websocket.TextMessage, resBytes)
		webClientsConnLock.Unlock()
	}

	return
}
