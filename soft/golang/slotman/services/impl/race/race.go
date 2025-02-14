package race

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/keyin"
	"slotman/services/iface/pilots"
	"slotman/services/iface/race"
	"slotman/services/iface/server"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/iface/tacho"
	"slotman/services/iface/teams"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"time"
)

type Service struct {
	srv server.Interface
	kin keyin.Interface
	amp ampel.Interface
	sdi speedi.Interface
	sdo speedo.Interface
	tco tacho.Interface
	tms teams.Interface
	plt pilots.Interface

	raceState     slotman.RaceState
	raceStateDone slotman.RaceState

	raceInfos []*slotman.RaceInfo

	trackStates   []slotman.TrackState
	trackVoltages []int

	setup *slotman.Setup

	roundsToGo int

	servicesReady bool
	looperStarted bool

	doExit bool
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.srv, err = server.GetInstance()
	if err != nil {
		return
	}

	singleTon.kin, err = keyin.GetInstance()
	if err != nil {
		return
	}

	singleTon.raceState = slotman.RaceStateIdle

	singleTon.trackStates = make([]slotman.TrackState, slotman.MaxTracks)
	singleTon.trackVoltages = make([]int, slotman.MaxTracks)

	singleTon.raceInfos = make([]*slotman.RaceInfo, slotman.MaxTracks)

	for track := range singleTon.raceInfos {
		singleTon.raceInfos[track] = &slotman.RaceInfo{
			What:  "info",
			Mode:  "set",
			Track: track,
		}
	}

	singleTon.kin.Subscribe(singleTon)

	singleTon.srv.Subscribe("tracks", singleTon)

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	singleTon.srv.Unsubscribe("tracks")

	singleTon.kin.Unsubscribe(singleTon)

	singleTon.doExit = true

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return race.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
