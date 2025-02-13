package race

import (
	"errors"
	"slotman/services/type/slotman"
	"slotman/utils/log"
)

func (sv *Service) GetRaceState() (state slotman.RaceState) {
	state = sv.raceState
	return
}

func (sv *Service) GetTrackStates() (trackStates []slotman.TrackState) {
	trackStates = sv.trackStates
	return
}

func (sv *Service) GetTrackVoltages() (trackVoltages []int) {
	trackVoltages = sv.trackVoltages
	return
}

func (sv *Service) GetRoundsToGo() (rounds int) {
	rounds = sv.roundsToGo
	return
}

func (sv *Service) GetRaceInfo(track int) (raceInfo *slotman.RaceInfo, err error) {

	if track < 0 || track >= slotman.MaxTracks {
		err = errors.New("bad track number")
		log.Cerror(err)
		return
	}

	raceInfo = sv.raceInfos[track]
	return
}

func (sv *Service) GetRaceInfos() (raceInfos []*slotman.RaceInfo) {
	raceInfos = sv.raceInfos
	return
}
