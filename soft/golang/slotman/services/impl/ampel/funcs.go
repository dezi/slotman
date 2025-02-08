package ampel

import "slotman/utils/log"

func (sv *Service) SetIdle() {

	log.Printf("######################## set idle")

	sv.ampelState = AmpelStateIdle
	go sv.patternIdle()
}
