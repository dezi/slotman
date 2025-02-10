package race

import "slotman/services/type/race"

func (sv *Service) GetRaceState() (state race.RaceState) {
	state = sv.raceState
	return
}

func (sv *Service) GetTracksVoltage() (tracksVoltage []int) {
	tracksVoltage = sv.tracksVoltage
	return
}
