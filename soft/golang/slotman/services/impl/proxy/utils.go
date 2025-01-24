package proxy

import (
	"errors"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) getTarget() (target string, err error) {

	hostName, err := os.Hostname()
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

	sv.webServer, _, err = websocket.DefaultDialer.Dial(wsUrl.String(), nil)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}
