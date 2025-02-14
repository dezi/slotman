package race

import (
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"time"
)

func (sv *Service) looper() {

	log.Printf("Race looper started...")
	defer log.Printf("Race looper done.")

	var waitingReady time.Time

	for !sv.doExit {

		time.Sleep(time.Millisecond * 40)

		if sv.raceState == slotman.RaceStateRaceWaiting {

			if sv.raceStateDone != slotman.RaceStateRaceWaiting {
				waitingReady = time.Now()
			}

			tracksReady := 0
			tracksActive := 0

			for _, ready := range sv.trackStates {

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

			if tracksActive > 0 && tracksActive == tracksReady &&
				time.Now().Unix()-waitingReady.Unix() > 3 {
				sv.raceState = slotman.RaceStateRaceStarting
				sv.sdo.SetTrackEnableAll(true)
			}
		}

		if sv.raceState == sv.raceStateDone {
			continue
		}

		sv.raceStateDone = sv.raceState

		switch sv.raceState {

		case slotman.RaceStateIdle:
			sv.sdo.SetTrackEnableAll(true)
			sv.amp.SetIdle()

		case slotman.RaceStateRaceStarting:
			sv.amp.SetRaceStart()

		case slotman.RaceStateRaceRunning:
			sv.amp.SetRaceRunning()

		case slotman.RaceStateRaceSuspended:
			sv.amp.SetRaceSuspended()

		case slotman.RaceStateRaceWaiting:
			sv.amp.SetRaceWaiting(sv.trackStates)

		case slotman.RaceStateRaceFinished:
			sv.amp.SetRaceFinished(sv.trackStates)
		}
	}
}
