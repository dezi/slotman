package race

import (
	"slotman/services/type/race"
	"slotman/utils/log"
)

func (sv *Service) OnAmpelClickShort() {
	log.Printf("OnAmpelClickShort...")

	if sv.raceState == race.RaceStateRaceRunning {
		sv.raceState = race.RaceStateRaceSuspended
		sv.amp.SetRaceSuspend()
		return
	}

	if sv.raceState == race.RaceStateRaceSuspended {
		sv.raceState = race.RaceStateRaceRunning
		sv.amp.SetRaceRestart()
		return
	}

}

func (sv *Service) OnAmpelClickLong() {
	log.Printf("OnAmpelClickLong...")

	if sv.raceState == race.RaceStateIdle {
		sv.raceState = race.RaceStateRaceStart
		return
	}

	sv.raceState = race.RaceStateIdle
}

func (sv *Service) OnMotoronVoltage(tracks []int, voltageMv uint32) {
	log.Printf("OnMotoronVoltage tracks=%v voltageMv=%d", tracks, voltageMv)
}

func (sv *Service) OnRaceStarted() {
	log.Printf("OnRaceStarted...")
	sv.raceState = race.RaceStateRaceRunning
}
