package race

import (
	"math/rand"
	"slotman/services/type/race"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"sort"
	"time"
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

		if sv.roundsToGo == 0 {
			sv.roundsToGo = 5
		}

		//
		// Fake pilots setup.
		//

		pilots := sv.plt.GetAllPilots()

		for shuffle := 0; shuffle < len(pilots); shuffle++ {

			inx1 := rand.Intn(len(pilots))
			inx2 := rand.Intn(len(pilots))

			pilot := pilots[inx1]
			pilots[inx1] = pilots[inx2]
			pilots[inx2] = pilot
		}

		for tracks := 0; tracks < slotman.MaxTracks; tracks++ {

			if tracks >= len(pilots) {
				continue
			}

			sv.raceRecords[tracks] = race.RaceRecord{
				Pilot: pilots[tracks],
			}
		}

		sv.raceState = race.RaceStateRaceWaiting
		return
	}

	sv.raceState = race.RaceStateIdle
	sv.roundsToGo = 0
}

func (sv *Service) OnMotoronVoltage(tracks []int, voltageMv uint32) {

	log.Printf("OnMotoronVoltage tracks=%v voltageMv=%d", tracks, voltageMv)

	for _, track := range tracks {
		sv.tracksVoltage[track] = int(voltageMv)
	}
}

func (sv *Service) OnRaceStarted() {

	log.Printf("OnRaceStarted...")

	sv.raceState = race.RaceStateRaceRunning

	for track := range sv.raceRecords {
		sv.raceRecords[track].LastRoundTime = time.Now()
	}
}

func (sv *Service) OnEnterStartPosition(track int) {

	log.Printf("OnEnterStartPosition track=%d", track)

	sv.tracksReady[track] = 2

	if sv.raceState == race.RaceStateRaceWaiting {
		sv.sdo.SetTrackEnable(track, false)
	}
}

func (sv *Service) OnLeaveStartPosition(track int) {
	log.Printf("OnLeaveStartPosition track=%d", track)
	sv.tracksReady[track] = 1
}

func (sv *Service) OnRoundCompleted(track int, roundMillis int) {

	if track < 0 || track >= slotman.MaxTracks {
		return
	}

	secs := float64(roundMillis) / 1000

	sv.raceRecords[track].Rounds++
	sv.raceRecords[track].ActRound = secs
	sv.raceRecords[track].LastRoundTime = time.Now()

	if sv.raceRecords[track].TopRound == 0 ||
		sv.raceRecords[track].TopRound > sv.raceRecords[track].ActRound {
		sv.raceRecords[track].TopRound = sv.raceRecords[track].ActRound
	}

	log.Printf("OnRoundCompleted     track=%d secs=%0.3f", track, secs)

	//
	// Re-score pilots order.
	//

	sortRecords := make([]*race.RaceRecord, slotman.MaxTracks)

	for inx := 0; inx < slotman.MaxTracks; inx++ {
		sortRecords[inx] = &sv.raceRecords[inx]
	}

	sort.Slice(sortRecords, func(i, j int) bool {
		if sortRecords[i].Rounds == sortRecords[j].Rounds {

			return sortRecords[i].LastRoundTime.Unix() < sortRecords[j].LastRoundTime.Unix()

		} else {

			return sortRecords[i].Rounds > sortRecords[j].Rounds
		}
	})

	for position, record := range sortRecords {
		record.Position = position + 1
	}
}

func (sv *Service) OnSpeedMeasurement(track int, speed float64) {

	if track < 0 || track >= slotman.MaxTracks {
		return
	}

	sv.raceRecords[track].Rounds++
	sv.raceRecords[track].ActSpeed = speed

	if sv.raceRecords[track].TopSpeed == 0 ||
		sv.raceRecords[track].TopSpeed < sv.raceRecords[track].ActSpeed {
		sv.raceRecords[track].TopSpeed = sv.raceRecords[track].ActSpeed
	}

	log.Printf("OnSpeedMeasurement   track=%d speed=%5.1f km/h", track, speed)
}

func (sv *Service) OnEmergencyStopNow(track int) {

	log.Printf("OnEmergencyStopNow track=%d", track)

	if sv.raceState != race.RaceStateRaceWaiting {
		return
	}

	log.Printf("OnEmergencyStopNow track=%d disable now", track)

	sv.sdo.SetTrackEnable(track, false)

	err := sv.sdo.SetSpeed(track, 0, true)
	log.Cerror(err)
}
