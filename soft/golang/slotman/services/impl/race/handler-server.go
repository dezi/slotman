package race

import (
	"encoding/json"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) OnRequestFromClient(appId simple.UUIDHex, what string, reqBytes []byte) {
	switch what {
	case "tracks":
		err := sv.handleTracks(appId, reqBytes)
		log.Cerror(err)
	}
}

func (sv *Service) handleTracks(appId simple.UUIDHex, reqBytes []byte) (err error) {

	tracks := slotman.Tracks{}
	err = json.Unmarshal(reqBytes, &tracks)
	if err != nil {
		log.Cerror(err)
		return
	}

	if tracks.Mode == "get" {

		tracks = sv.setup.Tracks

		reqBytes, err = simple.MarshalJsonClean(tracks)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = sv.srv.Transmit(appId, reqBytes)
		log.Cerror(err)

		return
	}

	if tracks.Mode == "set" {

		sv.setup.Tracks = tracks

		log.Printf("Tracks tracks=%d", sv.setup.Tracks.Tracks)

		err = sv.srv.Broadcast(reqBytes)
		log.Cerror(err)

		return
	}

	return
}
