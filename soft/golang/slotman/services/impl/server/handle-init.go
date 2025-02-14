package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"slotman/services/type/server"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) handleInit(appId simple.UUIDHex, ws *websocket.Conn, jsonBytes []byte) (err error) {

	init := server.Message{}
	err = json.Unmarshal(jsonBytes, &init)
	if err != nil {
		log.Cerror(err)
		return
	}

	if init.Mode == "get" {

		tracks := sv.setup.Tracks

		jsonBytes, err = simple.MarshalJsonClean(tracks)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = ws.WriteMessage(websocket.TextMessage, jsonBytes)
		log.Cerror(err)

		race := sv.setup.Race

		jsonBytes, err = simple.MarshalJsonClean(race)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = ws.WriteMessage(websocket.TextMessage, jsonBytes)
		log.Cerror(err)

		for _, pilot := range sv.setup.Pilots {

			pilot.What = "pilot"
			pilot.Mode = "set"

			jsonBytes, err = simple.MarshalJsonClean(pilot)
			if err != nil {
				log.Cerror(err)
				return
			}

			err = ws.WriteMessage(websocket.TextMessage, jsonBytes)
			log.Cerror(err)
		}

		return
	}

	return
}
