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

	res := Speedi{}

	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		log.Cerror(err)
		return
	}

	switch res.What {
	case SpeediWhatSpeed:
		_ = sv.handleLocalSpeed(res.Track, res.RawSpeed, nil)
	}

	return
}
