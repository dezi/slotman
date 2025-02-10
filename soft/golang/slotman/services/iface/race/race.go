package race

import (
	"slotman/services/impl/provider"
	"slotman/services/type/race"
)

const (
	Service provider.Service = "serviceRace"
)

type Interface interface {
	GetName() (name provider.Service)

	GetRaceState() (state race.RaceState)
	GetTracksReady() (tracksReady []int)
	GetTracksVoltage() (tracksVoltage []int)
	GetRoundsToGo() (rounds int)

	OnAmpelClickShort()
	OnAmpelClickLong()

	OnRaceStarted()

	OnMotoronVoltage(tracks []int, voltageMv uint32)

	OnEnterStartPosition(track int)
	OnLeaveStartPosition(track int)
	OnRoundCompleted(track int, roundMillis int)
	OnSpeedMeasurement(track int, speed float64)
	OnEmergencyStopNow(track int)
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
