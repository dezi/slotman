package tacho

import (
	"slotman/utils/log"
	"time"
)

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
	trackState.RoundTs = time.Now()
	sv.trackStates[track] = trackState

	sv.mapsLock.Unlock()
}

func (sv *Service) OnRoundCompleted(track int) {

	now := time.Now()

	sv.mapsLock.Lock()

	trackState := sv.trackStates[track]

	trackState.RoundMillis = int(now.UnixMilli() - trackState.RoundTs.UnixMilli())
	trackState.RoundTs = time.Now()

	sv.trackStates[track] = trackState

	sv.mapsLock.Unlock()

	log.Printf("OnRoundCompleted     track=%d secs=%0.3f",
		track, float64(trackState.RoundMillis)/1000)
}

func (sv *Service) OnSpeedMeasurement(track int, speed float64) {

	log.Printf("OnSpeedMeasurement   track=%d speed=%5.1f km/h", track, speed)

	sv.mapsLock.Lock()

	trackState := sv.trackStates[track]
	trackState.SpeedKmh = speed
	trackState.SpeedTS = time.Now()
	sv.trackStates[track] = trackState

	sv.mapsLock.Unlock()
}

func (sv *Service) OnEmergencyStopNow(track int) {
	log.Printf("OnEmergencyStopNow   track=%d", track)
}
