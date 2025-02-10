package race

import (
	"errors"
	"slotman/services/type/race"
	"slotman/services/type/slotman"
	"slotman/utils/log"
)

func (sv *Service) GetRaceState() (state race.RaceState) {
	state = sv.raceState
	return
}

func (sv *Service) GetTracksReady() (tracksReady []int) {
	tracksReady = sv.tracksReady
	return
}

func (sv *Service) GetTracksVoltage() (tracksVoltage []int) {
	tracksVoltage = sv.tracksVoltage
	return
}

func (sv *Service) GetRoundsToGo() (rounds int) {
	rounds = sv.roundsToGo
	return
}

func (sv *Service) GetRaceRecord(track int) (raceRecord race.RaceRecord, err error) {

	if track < 0 || track >= slotman.MaxTracks {
		err = errors.New("bad track number")
		log.Cerror(err)
		return
	}

	raceRecord = sv.raceRecords[track]
	return
}

func (sv *Service) GetRaceRecords() (raceRecords []race.RaceRecord) {
	raceRecords = sv.raceRecords
	return
}
