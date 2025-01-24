package proxy

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"slotman/services/type/proxy"
	"slotman/utils/log"
	"strings"
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

	sender := r.RemoteAddr

	log.Printf("Started websocket sender=%s...", sender)
	defer log.Printf("Stopped websocket sender=%s.", sender)

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

	sv.mapsLock.Lock()

	if sv.webSockets[sender] != nil {
		_ = sv.webSockets[sender].Close()
	}

	sv.webSockets[sender] = ws

	sv.mapsLock.Unlock()

	var mType int
	var tryErr error
	var jsonBytes []byte

	for {

		mType, jsonBytes, tryErr = ws.ReadMessage()
		if tryErr != nil {
			break
		}

		if mType != websocket.TextMessage {
			continue
		}

		//log.Printf("Message mType=%d jsonBytes=%sd", mType, string(jsonBytes))

		message := proxy.Message{}
		err = json.Unmarshal(jsonBytes, &message)
		if err != nil {
			log.Cerror(err)
			break
		}

		switch message.What {
		//case "race":
		//	err = sv.handleRace(appId, ws, jsonBytes)
		//case "pilot":
		//	err = sv.handlePilot(appId, ws, jsonBytes)
		//case "tracks":
		//	err = sv.handleTracks(appId, ws, jsonBytes)
		//case "controller":
		//	err = sv.handleController(appId, ws, jsonBytes)
		}
	}

	sv.mapsLock.Lock()

	if sv.webSockets[sender] == ws {
		_ = ws.Close()
		delete(sv.webSockets, sender)
	}

	sv.mapsLock.Unlock()
}
