package race

import (
	"slotman/services/type/race"
	"slotman/utils/log"
	"time"
)

func (sv *Service) looper() {

	log.Printf("Race looper started...")
	defer log.Printf("Race looper done.")

	var waitingReady time.Time

	for !sv.doExit {

		time.Sleep(time.Millisecond * 40)

		if sv.raceState == race.RaceStateRaceWaiting {

			if sv.raceStateDone != race.RaceStateRaceWaiting {
				waitingReady = time.Now()
			}

			tracksReady := 0
			tracksActive := 0

			for _, ready := range sv.tracksReady {

				if ready == 0 {
					continue
				}

				tracksActive++

				if ready == 1 {
					waitingReady = time.Now()
					continue
				}

				tracksReady++
			}

			if tracksActive == tracksReady && time.Now().Unix()-waitingReady.Unix() > 3 {
				sv.raceState = race.RaceStateRaceStarting
			}
		}

		if sv.raceState == sv.raceStateDone {
			continue
		}

		sv.raceStateDone = sv.raceState

		switch sv.raceState {

		case race.RaceStateIdle:
			sv.amp.SetIdle()

		case race.RaceStateRaceStarting:
			sv.amp.SetRaceStart()

		case race.RaceStateRaceRunning:
			sv.amp.SetRaceRunning()

		case race.RaceStateRaceSuspended:
			sv.amp.SetRaceSuspended()

		case race.RaceStateRaceWaiting:
			sv.amp.SetRaceWaiting(sv.tracksReady)
		}
	}
}
