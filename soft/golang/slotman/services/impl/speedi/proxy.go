package speedi

import (
	"encoding/json"
	"slotman/utils/log"
)

func (sv *Service) OnMessageFromClient(reqBytes []byte) (resBytes []byte, err error) {

	req := Speedi{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	switch req.What {
	case SpeediWhatOpen:
		// initialize and open speed reader.
	case SpeediWhatClose:
		// close speed reader.
	}

	return
}

func (sv *Service) OnMessageFromServer(resBytes []byte) {

}
