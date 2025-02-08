package race

import (
	"slotman/services/type/race"
	"slotman/utils/log"
	"time"
)

func (sv *Service) looper() {

	log.Printf("Race looper started...")
	defer log.Printf("Race looper done.")

	for !sv.doExit {

		time.Sleep(time.Millisecond * 40)

		if sv.raceState == sv.raceStateDone {
			continue
		}

		sv.raceStateDone = sv.raceState

		switch sv.raceState {
		case race.RaceStateIdle:
			sv.amp.SetIdle()
		}
	}
}
