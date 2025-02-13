package ampel

import "slotman/services/type/slotman"

func (sv *Service) SetRoundsToGo(roundsToGo int) {
	sv.roundsToGo = roundsToGo
}

func (sv *Service) SetIdle() {
	sv.ampelState = AmpelStateIdle
	go sv.patternIdle()
}

func (sv *Service) SetRaceStart() {
	sv.ampelState = AmpelStateRaceStart
	go sv.patternRaceStart()
}

func (sv *Service) SetRaceWaiting(trackStates []slotman.TrackState) {

	sv.waitingTracksReady = trackStates

	if sv.ampelState != AmpelStateRaceWaiting {
		sv.ampelState = AmpelStateRaceWaiting
		go sv.patternRaceWaiting()
	}
}

func (sv *Service) SetRaceSuspended() {
	sv.ampelState = AmpelStateRaceSuspended
	go sv.patternRaceSuspended()
}

func (sv *Service) SetRaceRunning() {
	sv.ampelState = AmpelStateRaceRunning
	go sv.patternRaceRunning()
}
