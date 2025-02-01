package speedi

import (
	"encoding/json"
	"slotman/utils/log"
)

func (sv *Service) OnMessageFromClient(reqBytes []byte) (resBytes []byte, err error) {
	_ = reqBytes
	return
}

func (sv *Service) OnMessageFromServer(resBytes []byte) {

	var err error

	req := Speedi{}

	err = json.Unmarshal(resBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	switch req.What {
	case SpeediWhatSpeed:
	}

	return
}
