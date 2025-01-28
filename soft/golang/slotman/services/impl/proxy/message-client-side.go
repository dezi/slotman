package proxy

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"slotman/services/type/proxy"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) proxyRequest(req proxy.MessageIface) (res []byte, err error) {

	sv.webServerConnLock.Lock()
	defer sv.webServerConnLock.Unlock()

	if sv.webServerConn == nil {
		err = sv.createConnect()
		if err != nil {
			return
		}
	}

	uuid := simple.NewUuidHex()
	resc := make(chan []byte, 1)

	sv.webServerChanLock.Lock()
	sv.webServerChan[uuid] = resc
	sv.webServerChanLock.Unlock()

	req.SetUuid(uuid)

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return
	}

	log.Printf("proxyRequest req=%s", string(reqBytes))

	err = sv.webServerConn.WriteMessage(websocket.TextMessage, reqBytes)
	if err != nil {
		log.Cerror(err)
		_ = sv.webServerConn.Close()
		sv.webServerConn = nil
		return
	}

	res = <-resc
	close(resc)

	sv.webServerChanLock.Lock()
	delete(sv.webServerChan, uuid)
	sv.webServerChanLock.Unlock()

	//var mType int
	//mType, res, err = sv.webServerConn.ReadMessage()
	//if mType != websocket.TextMessage {
	//	err = errors.New("wrong message type received")
	//	log.Cerror(err)
	//	return
	//}

	//log.Printf("proxyRequest res=%s", string(res))

	return
}

func (sv *Service) createConnect() (err error) {

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

	go sv.connectReadLoop()

	log.Printf("Connected wsUrl=%s", wsUrl.String())
	return
}

func (sv *Service) connectReadLoop() {

	var mType int
	var res []byte
	var err error

	for sv.webServerConn != nil {

		mType, res, err = sv.webServerConn.ReadMessage()
		if err != nil {
			log.Cerror(err)
			return
		}

		if mType != websocket.TextMessage {
			err = errors.New("wrong message type received")
			log.Cerror(err)
			continue
		}

		log.Printf("proxyRequest res=%s", string(res))

		message := &proxy.Message{}
		err = json.Unmarshal(res, message)
		if err != nil {
			log.Cerror(err)
			continue
		}

		uuid := message.Uuid

		sv.webServerChanLock.Lock()
		resc := sv.webServerChan[uuid]
		sv.webServerChanLock.Unlock()

		if resc != nil {
			resc <- res
			continue
		}

		log.Printf("########### out of band uuid=%s...", uuid)
		//
		// Handle out of band push message.
		//
	}
}
