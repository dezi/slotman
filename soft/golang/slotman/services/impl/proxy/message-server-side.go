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

	addrParts := strings.Split(r.RemoteAddr, ":")
	if len(addrParts) < 2 {
		http.NotFound(w, r)
		return
	}

	sender := strings.Join(addrParts[:len(addrParts)-1], ":")

	log.Printf("Started websocket remoteAddr=%s...", r.RemoteAddr)
	defer log.Printf("Stopped websocket remoteAddr=%s.", r.RemoteAddr)

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

	sv.deleteClientConnect(sender)

	sv.webClientsLock.Lock()
	sv.webClientsConns[sender] = ws
	sv.webClientsLock.Unlock()

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

		//log.Printf("Recv reqBytes=%s", string(reqBytes))

		message := proxy.Message{}
		err = json.Unmarshal(reqBytes, &message)
		if err != nil {
			log.Cerror(err)
			break
		}

		var resBytes []byte

		switch message.Area {
		case proxy.AreaGpio:
			resBytes, err = sv.handleGpio(sender, reqBytes)
		case proxy.AreaI2c:
			resBytes, err = sv.handleI2c(sender, reqBytes)
		case proxy.AreaSpi:
			resBytes, err = sv.handleSpi(sender, reqBytes)
		case proxy.AreaUart:
			resBytes, err = sv.handleUart(sender, reqBytes)
		}

		if err != nil {
			log.Cerror(err)
			break
		}

		err = ws.WriteMessage(websocket.TextMessage, resBytes)
		if err != nil {
			log.Cerror(err)
			break
		}
	}

	sv.deleteClientConnect(sender)
}

func (sv *Service) deleteClientConnect(sender string) {

	sv.webClientsLock.Lock()
	defer sv.webClientsLock.Unlock()

	if sv.webClientsConns[sender] == nil {
		return
	}

	log.Printf("Delete socket sender=%s", sender)

	_ = sv.webClientsConns[sender].Close()
	delete(sv.webClientsConns, sender)

	sv.gpioDevLock.Lock()

	for oldSender, gpioDev := range sv.gpioDevMap {
		if strings.HasPrefix(oldSender, sender) {
			log.Printf("Delete GPIO sender=%s pin=%d", sender, gpioDev.GetPinNo())
			_ = gpioDev.Close()
			sv.gpioDevMap[oldSender] = nil
		}
	}

	sv.gpioDevLock.Unlock()

	sv.spiDevLock.Lock()

	for oldSender, spiDev := range sv.spiDevMap {
		if strings.HasPrefix(oldSender, sender) {
			log.Printf("Delete SPI  sender=%s dev=%s", sender, spiDev.GetDevice())
			_ = spiDev.Close()
			sv.spiDevMap[oldSender] = nil
		}
	}

	sv.spiDevLock.Unlock()

	sv.uartDevLock.Lock()

	for oldSender, uartDev := range sv.uartDevMap {
		if strings.HasPrefix(oldSender, sender) {
			log.Printf("Delete UART sender=%s dev=%s", sender, uartDev.GetDevice())
			_ = uartDev.Close()
			sv.uartDevMap[oldSender] = nil
		}
	}

	sv.uartDevLock.Unlock()
}
