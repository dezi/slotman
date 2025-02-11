package server

import (
	"github.com/gorilla/websocket"
)

func (sv *Service) Broadcast(resBytes []byte) (err error) {

	sv.webClientsLock.Lock()
	defer sv.webClientsLock.Unlock()

	for sender, webClientsConn := range sv.webClientsConns {

		webClientsConnLock := sv.webClientsConnsLocks[sender]

		if webClientsConn == nil || webClientsConnLock == nil {
			continue
		}

		webClientsConnLock.Lock()
		_ = webClientsConn.WriteMessage(websocket.TextMessage, resBytes)
		webClientsConnLock.Unlock()
	}

	return
}
