package race

import (
	"encoding/json"
	"math/rand"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"sort"
	"time"
)

func (sv *Service) OnAmpelClickShort() {

	log.Printf("OnAmpelClickShort...")

	if sv.raceState == slotman.RaceStateIdle {

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

	if sv.raceState == slotman.RaceStateRaceRunning {
		sv.raceState = slotman.RaceStateRaceSuspended
		return
	}

	if sv.raceState == slotman.RaceStateRaceSuspended {
		sv.raceState = slotman.RaceStateRaceRunning
		return
	}
}

func (sv *Service) OnAmpelClickLong() {

	log.Printf("OnAmpelClickLong...")

	if sv.raceState == slotman.RaceStateIdle {

		if sv.roundsToGo == 0 {
			sv.roundsToGo = 50
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

			sv.raceInfos[tracks].Rounds = 0
			sv.raceInfos[tracks].Position = 0
			sv.raceInfos[tracks].ActRound = 0
			sv.raceInfos[tracks].TopRound = 0
			sv.raceInfos[tracks].ActSpeed = 0
			sv.raceInfos[tracks].TopSpeed = 0
			sv.raceInfos[tracks].PilotUuid = pilots[tracks].Uuid
		}

		sv.raceState = slotman.RaceStateRaceWaiting

		sv.sdo.SetTrackFixedSpeed(0, 44)
		sv.sdo.SetTrackFixedSpeed(1, 44)

		return
	}

	sv.raceState = slotman.RaceStateIdle
	sv.roundsToGo = 0
}

func (sv *Service) OnMotoronVoltage(tracks []int, voltageMv uint32) {

	//log.Printf("OnMotoronVoltage tracks=%v voltageMv=%d", tracks, voltageMv)

	for _, track := range tracks {
		sv.tracksVoltage[track] = int(voltageMv)
	}
}

func (sv *Service) OnRaceStarted() {

	log.Printf("OnRaceStarted...")

	sv.raceState = slotman.RaceStateRaceRunning

	for track := range sv.raceInfos {
		sv.raceInfos[track].LastRoundTime = time.Now()
	}
}

func (sv *Service) OnEnterStartPosition(track int) {

	log.Printf("OnEnterStartPosition track=%d", track)

	sv.tracksReady[track] = 2

	if sv.raceState == slotman.RaceStateRaceWaiting {
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

	sv.raceInfos[track].Rounds++
	sv.raceInfos[track].ActRound = secs
	sv.raceInfos[track].LastRoundTime = time.Now()

	if sv.raceInfos[track].TopRound == 0 ||
		sv.raceInfos[track].TopRound > sv.raceInfos[track].ActRound {
		sv.raceInfos[track].TopRound = sv.raceInfos[track].ActRound
	}

	log.Printf("OnRoundCompleted     track=%d rounds=%d secs=%0.3f",
		track, sv.raceInfos[track].Rounds, secs)

	//
	// Re-score pilots order.
	//

	sortRecords := make([]*slotman.RaceInfo, slotman.MaxTracks)

	for inx := 0; inx < slotman.MaxTracks; inx++ {
		sortRecords[inx] = sv.raceInfos[inx]
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

	for ti, info := range sv.raceInfos {

		if sv.tracksReady[ti] == 0 {
			continue
		}

		resBytes, err := json.Marshal(info)
		if err != nil {
			log.Cerror(err)
			continue
		}

		err = sv.srv.Broadcast(resBytes)
		log.Cerror(err)
	}
}

func (sv *Service) OnSpeedMeasurement(track int, speed float64) {

	if track < 0 || track >= slotman.MaxTracks {
		return
	}

	sv.raceInfos[track].ActSpeed = speed

	if sv.raceInfos[track].TopSpeed == 0 ||
		sv.raceInfos[track].TopSpeed < sv.raceInfos[track].ActSpeed {
		sv.raceInfos[track].TopSpeed = sv.raceInfos[track].ActSpeed
	}

	log.Printf("OnSpeedMeasurement   track=%d speed=%5.1f km/h", track, speed)
}

func (sv *Service) OnEmergencyStopNow(track int) {

	log.Printf("OnEmergencyStopNow track=%d", track)

	if sv.raceState == slotman.RaceStateRaceWaiting {

		log.Printf("OnEmergencyStopNow track=%d disable now", track)

		sv.sdo.SetTrackEnable(track, false)

		err := sv.sdo.SetSpeed(track, 0, true)
		log.Cerror(err)

		return
	}

	if sv.raceState == slotman.RaceStateRaceRunning {

		if sv.raceInfos[track].Rounds+1 >= sv.roundsToGo {

			log.Printf("OnEmergencyStopNow track=%d finished now", track)

			sv.sdo.SetTrackEnable(track, false)

			err := sv.sdo.SetSpeed(track, 0, true)
			log.Cerror(err)
		}
	}
}
