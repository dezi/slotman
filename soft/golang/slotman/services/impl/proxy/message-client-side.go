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

func (sv *Service) CheckTarget() (ok bool) {

	var hostName string
	hostName, err := os.Hostname()
	if err != nil {
		return
	}

	_, ok = proxy.ProxyTargets[hostName]

	return
}

func (sv *Service) ProxyRequest(req proxy.Message) (res []byte, err error) {

	sv.webServerConnLock.Lock()

	if sv.webServerConn == nil {
		err = sv.createConnect()
		if err != nil {
			sv.webServerConnLock.Unlock()
			return
		}
	}

	sv.webServerConnLock.Unlock()

	resc := make(chan []byte, 1)
	defer close(resc)

	uuid := simple.NewUuidHex()

	sv.webServerChanLock.Lock()
	sv.webServerChan[uuid] = resc
	sv.webServerChanLock.Unlock()

	defer func() {
		sv.webServerChanLock.Lock()
		delete(sv.webServerChan, uuid)
		sv.webServerChanLock.Unlock()
	}()

	req.SetUuid(uuid)

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return
	}

	sv.webServerConnLock.Lock()
	defer sv.webServerConnLock.Unlock()

	if sv.webServerConn == nil {
		err = errors.New("no socket connect")
		log.Cerror(err)
		return
	}

	err = sv.webServerConn.WriteMessage(websocket.TextMessage, reqBytes)
	if err != nil {
		log.Cerror(err)
		_ = sv.webServerConn.Close()
		sv.webServerConn = nil
		return
	}

	res = <-resc

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

	var err error
	var conn *websocket.Conn
	var mType int
	var resBytes []byte

	for {

		conn = sv.webServerConn
		if conn == nil {
			break
		}

		mType, resBytes, err = conn.ReadMessage()
		if err != nil {
			log.Cerror(err)
			return
		}

		if mType != websocket.TextMessage {
			err = errors.New("wrong message type received")
			log.Cerror(err)
			continue
		}

		//log.Printf("ProxyRequest res=%s", string(res))

		msg := &message{}
		err = json.Unmarshal(resBytes, msg)
		if err != nil {
			log.Cerror(err)
			continue
		}

		uuid := msg.Uuid

		sv.webServerChanLock.Lock()
		resc := sv.webServerChan[uuid]
		sv.webServerChanLock.Unlock()

		if resc != nil {
			resc <- resBytes
			continue
		}

		//
		// Handle out of band push message.
		//

		//log.Printf("ProxyRequest OOB res=%s", string(resBytes))

		sv.subscribersLock.Lock()
		subscriber := sv.subscribers[msg.Area]
		sv.subscribersLock.Unlock()

		if subscriber != nil {
			go subscriber.OnMessageFromServer(resBytes)
		}
	}
}
