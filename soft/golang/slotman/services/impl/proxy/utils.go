package proxy

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) WriteMessage(message interface{}) (err error) {

	sv.webServerLock.Lock()
	defer sv.webServerLock.Unlock()

	if sv.webServerConn == nil {

		var hostName string
		hostName, err = os.Hostname()
		if err != nil {
			return
		}

		target, ok := proxy.ProxyTargets[hostName]
		if !ok {
			err = errors.New("no proxy target for host")
			return
		}

		wsUrl := url.URL{Scheme: "ws", Host: target, Path: "/ws"}
		log.Printf("Connecting wsUrl=%s", wsUrl.String())

		sv.webServerConn, _, err = websocket.DefaultDialer.Dial(wsUrl.String(), nil)
		if err != nil {
			log.Cerror(err)
			return
		}

		log.Printf("Connected wsUrl=%s", wsUrl.String())
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return
	}

	log.Printf("########### WriteMessage messageBytes=%s", string(messageBytes))

	err = sv.webServerConn.WriteMessage(websocket.TextMessage, messageBytes)
	if err == nil {
		return
	}

	log.Cerror(err)

	_ = sv.webServerConn.Close()
	sv.webServerConn = nil

	return
}
