package turner

import (
	"encoding/json"
	"slotman/utils/log"
)

func (sv *Service) OnMessageFromClient(reqBytes []byte) (resBytes []byte, err error) {

	log.Printf("################# Turner OnMessageFromClient...")

	req := Turner{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	switch req.What {
	case TurnerWhatBlip:

		log.Printf("################# size=%d", len(req.BlipImage))
		req.BlipImage = nil
	}

	resBytes, err = json.Marshal(req)
	log.Cerror(err)

	return
}

func (sv *Service) OnMessageFromServer(resBytes []byte) {
	_ = resBytes
	return
}
