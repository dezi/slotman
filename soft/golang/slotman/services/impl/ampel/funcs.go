package ampel

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

func (sv *Service) SetRaceSuspend() {
	sv.ampelState = AmpelStateRaceSuspend
	go sv.patternRaceSuspend()
}

func (sv *Service) SetRaceRestart() {
	sv.ampelState = AmpelStateRaceRestart
	go sv.patternRaceRestart()
}
