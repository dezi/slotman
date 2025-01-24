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

	if !strings.HasPrefix(r.URL.String(), "/ws") {
		http.NotFound(w, r)
		return
	}

	if strings.Contains(r.URL.String(), "/..") {
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

	if sv.webClients[sender] != nil {
		_ = sv.webClients[sender].Close()
	}

	sv.webClients[sender] = ws

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

		log.Printf("Recv reqBytes=%s", string(reqBytes))

		message := proxy.Message{}
		err = json.Unmarshal(reqBytes, &message)
		if err != nil {
			log.Cerror(err)
			break
		}

		var resBytes []byte

		switch message.Area {
		case proxy.AreaGpio:
			resBytes, err = sv.handleGpio(reqBytes)
		}

		if err != nil {
			log.Cerror(err)
			break
		}

		log.Printf("Send resBytes=%s", string(resBytes))

		err = ws.WriteMessage(websocket.TextMessage, resBytes)
		if err != nil {
			log.Cerror(err)
			break
		}
	}

	sv.mapsLock.Lock()

	if sv.webClients[sender] == ws {
		_ = ws.Close()
		delete(sv.webClients, sender)
	}

	sv.mapsLock.Unlock()
}
