package tacho

import (
	"slotman/utils/log"
	"time"
)

func (sv *Service) OnSpeedMeasurement(track int, speed float64) {

	log.Printf("OnSpeedMeasurement track=%d speed=%5.1f", track, speed)

	sv.mapsLock.Lock()

	trackState := sv.trackStates[track]
	trackState.Speed = speed
	trackState.SpeedTS = time.Now()
	sv.trackStates[track] = trackState

	sv.mapsLock.Unlock()
}

func (sv *Service) OnEnterStartPosition(track int) {

	log.Printf("OnEnterStartPosition track=%d", track)

	sv.mapsLock.Lock()

	trackState := sv.trackStates[track]
	trackState.IsAtStart = true
	sv.trackStates[track] = trackState

	sv.mapsLock.Unlock()
}

func (sv *Service) OnLeaveStartPosition(track int) {

	log.Printf("OnLeaveStartPosition track=%d", track)

	sv.mapsLock.Lock()

	trackState := sv.trackStates[track]
	trackState.IsAtStart = false
	sv.trackStates[track] = trackState

	sv.mapsLock.Unlock()
}
