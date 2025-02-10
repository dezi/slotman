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

	if sv.raceState == race.RaceStateRaceWaiting {
		sv.raceState = race.RaceStateRaceStarting
	}

	if sv.raceState == race.RaceStateRaceRunning {
		sv.raceState = race.RaceStateRaceSuspended
		return
	}

	if sv.raceState == race.RaceStateRaceSuspended {
		sv.raceState = race.RaceStateRaceRunning
		return
	}
}

func (sv *Service) OnAmpelClickLong() {

	log.Printf("OnAmpelClickLong...")

	if sv.raceState == race.RaceStateIdle {
		sv.raceState = race.RaceStateRaceWaiting
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

func (sv *Service) OnEnterStartPosition(track int) {
	log.Printf("OnEnterStartPosition track=%d", track)
	sv.tracksReady[track] = 2
}

func (sv *Service) OnLeaveStartPosition(track int) {
	log.Printf("OnLeaveStartPosition track=%d", track)
	sv.tracksReady[track] = 1
}

func (sv *Service) OnRoundCompleted(track int, roundMillis int) {
	log.Printf("OnRoundCompleted     track=%d secs=%0.3f", track, float64(roundMillis)/1000)
}

func (sv *Service) OnSpeedMeasurement(track int, speed float64) {
	log.Printf("OnSpeedMeasurement   track=%d speed=%5.1f km/h", track, speed)
}

func (sv *Service) OnEmergencyStopNow(track int) {
	log.Printf("OnEmergencyStopNow   track=%d", track)
}
