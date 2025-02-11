package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) handleRace(appId simple.UUIDHex, ws *websocket.Conn, jsonBytes []byte) (err error) {

	race := slotman.Race{}
	err = json.Unmarshal(jsonBytes, &race)
	if err != nil {
		log.Cerror(err)
		return
	}

	if race.Mode == "get" {

		race = sv.setup.Race

		jsonBytes, err = simple.MarshalJsonClean(race)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = ws.WriteMessage(websocket.TextMessage, jsonBytes)
		log.Cerror(err)

		return
	}

	if race.Mode == "set" {

		sv.setup.Race = race

		log.Printf("Race title=%s tracks=%d rounds=%d",
			sv.setup.Race.Title, sv.setup.Race.Tracks, sv.setup.Race.Rounds)

		sv.broadcast(appId, jsonBytes)

		return
	}

	return
}
