package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) handleTracks(appId simple.UUIDHex, ws *websocket.Conn, jsonBytes []byte) (err error) {

	tracks := slotman.Tracks{}
	err = json.Unmarshal(jsonBytes, &tracks)
	if err != nil {
		log.Cerror(err)
		return
	}

	if tracks.Mode == "get" {

		tracks = sv.setup.Tracks

		jsonBytes, err = simple.MarshalJsonClean(tracks)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = ws.WriteMessage(websocket.TextMessage, jsonBytes)
		log.Cerror(err)

		return
	}

	if tracks.Mode == "set" {

		sv.setup.Tracks = tracks

		log.Printf("Tracks tracks=%d", sv.setup.Tracks.Tracks)

		sv.broadcast(appId, jsonBytes)

		return
	}

	return
}
