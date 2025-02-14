package ampel

import (
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
)

const (
	Service provider.Service = "serviceAmpel"
)

type Interface interface {
	GetName() (name provider.Service)

	SetRoundsToGo(roundsToGo int)

	SetIdle()

	SetRaceStart()
	SetRaceRunning()
	SetRaceSuspended()

	SetRaceWaiting(trackStates []slotman.TrackState)
	SetRaceFinished(trackStates []slotman.TrackState)
}

func GetInstance() (iface Interface, err error) {

	baseProvider, err := provider.GetProvider(Service)
	if err != nil {
		return
	}

	iface = baseProvider.(Interface)
	if iface == nil {
		err = provider.ErrNotFound(Service)
		return
	}

	return
}
