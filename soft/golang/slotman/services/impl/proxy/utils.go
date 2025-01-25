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

func (sv *Service) proxyRequest(req interface{}) (res []byte, err error) {

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

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return
	}

	//log.Printf("proxyRequest req=%s", string(reqBytes))

	err = sv.webServerConn.WriteMessage(websocket.TextMessage, reqBytes)
	if err != nil {
		log.Cerror(err)
		_ = sv.webServerConn.Close()
		sv.webServerConn = nil
		return
	}

	var mType int
	mType, res, err = sv.webServerConn.ReadMessage()
	if mType != websocket.TextMessage {
		err = errors.New("wrong message type received")
		log.Cerror(err)
		return
	}

	//log.Printf("proxyRequest res=%s", string(res))

	return
}
