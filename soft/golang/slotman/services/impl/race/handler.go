package race

import (
	"slotman/services/type/race"
	"slotman/utils/log"
)

func (sv *Service) OnAmpelClickShort() {

	log.Printf("OnAmpelClickShort...")

	if sv.raceState == race.RaceStateIdle {

		switch sv.roundsToGo {
		case 0:
			sv.roundsToGo = 5
		case 5:
			sv.roundsToGo = 10
		case 10:
			sv.roundsToGo = 25
		case 25:
			sv.roundsToGo = 50
		case 50:
			sv.roundsToGo = 100
		case 100:
			sv.roundsToGo = 0
		}

		sv.amp.SetRoundsToGo(sv.roundsToGo)
	}

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
	sv.roundsToGo = 0
}

func (sv *Service) OnMotoronVoltage(tracks []int, voltageMv uint32) {
	log.Printf("OnMotoronVoltage tracks=%v voltageMv=%d", tracks, voltageMv)
}

func (sv *Service) OnRaceStarted() {
	log.Printf("OnRaceStarted...")
	sv.raceState = race.RaceStateRaceRunning
}
