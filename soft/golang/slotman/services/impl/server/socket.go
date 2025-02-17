package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"slotman/services/type/server"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
	"sync"
)

func (sv *Service) handleWs(w http.ResponseWriter, r *http.Request) {

	if !strings.HasPrefix(r.URL.String(), "/ws/") {
		http.NotFound(w, r)
		return
	}

	if strings.Contains(r.URL.String(), "/..") {
		http.NotFound(w, r)
		return
	}

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	appId := simple.UUIDHex(parts[2])

	log.Printf("Started websocket appId=%s...", appId)
	defer log.Printf("Stopped websocket appId=%s.", appId)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Cerror(err)
		return
	}

	var wsLock sync.Mutex

	sv.mapsLock.Lock()

	if sv.webClientsConns[appId] != nil {
		_ = sv.webClientsConns[appId].Close()
	}

	sv.webClientsConns[appId] = ws
	sv.webClientsConnsLocks[appId] = &wsLock

	sv.mapsLock.Unlock()

	var mType int
	var tryErr error
	var reqBytes []byte

	for {

		mType, reqBytes, tryErr = ws.ReadMessage()
		if tryErr != nil {
			break
		}

		if mType != websocket.TextMessage {
			continue
		}

		//log.Printf("Message mType=%d reqBytes=%sd", mType, string(reqBytes))

		message := server.Message{}
		err = json.Unmarshal(reqBytes, &message)
		if err != nil {
			log.Cerror(err)
			break
		}

		sv.subscribersLock.Lock()
		subscriber := sv.subscribers[message.What]
		sv.subscribersLock.Unlock()

		if subscriber != nil {
			go subscriber.OnRequestFromClient(appId, message.What, reqBytes)
		}
	}

	sv.mapsLock.Lock()

	if sv.webClientsConns[appId] == ws {
		_ = ws.Close()
		delete(sv.webClientsConns, appId)
	}

	sv.mapsLock.Unlock()
}

func (sv *Service) broadcast(appId simple.UUIDHex, message []byte) {

	sv.mapsLock.Lock()

	log.Printf("########## bcast len=%d", len(sv.webClientsConns))

	for appIdOld, ws := range sv.webClientsConns {

		if appIdOld == appId {
			continue
		}

		log.Printf("###### bcast message=%s", string(message))

		_ = ws.WriteMessage(websocket.TextMessage, message)
	}

	sv.mapsLock.Unlock()

	return
}
