package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) handlePilot(appId simple.UUIDHex, ws *websocket.Conn, jsonBytes []byte) (err error) {

	pilot := &slotman.Pilot{}
	err = json.Unmarshal(jsonBytes, pilot)
	if err != nil {
		log.Cerror(err)
		return
	}

	if pilot.Mode == "set" {

		sv.mapsLock.Lock()
		sv.setup.Pilots[pilot.AppUuid] = pilot
		sv.mapsLock.Unlock()

		log.Printf("Pilot first=%s last=%s car=%s",
			pilot.FirstName, pilot.LastName, pilot.Car)

		sv.broadcast("", jsonBytes)

		return
	}

	return
}
