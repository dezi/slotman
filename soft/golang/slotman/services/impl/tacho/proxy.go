package tacho

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

	res := Tacho{}

	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		log.Cerror(err)
		return
	}

	switch res.What {
	case TachoWhatTacho:
		state := TachoState{
			active: res.Active,
			time:   res.Time,
		}

		sv.mapsLock.Lock()
		sv.tachoStates[res.Pin] = state
		sv.mapsLock.Unlock()

		sv.handleLocalTacho(res.Pin, state)
	}

	return
}
