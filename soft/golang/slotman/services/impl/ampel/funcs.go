package ampel

func (sv *Service) SetIdle() {
	sv.ampelState = AmpelStateIdle
	go sv.patternIdle()
}
