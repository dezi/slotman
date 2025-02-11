package speedo

import (
	"slotman/services/impl/provider"
)

const (
	Service provider.Service = "serviceSpeedo"
)

type Interface interface {
	GetName() (name provider.Service)

	SetSpeed(track int, percent float64, now bool) (err error)
	SetTrackEnable(track int, enable bool)
	SetTrackEnableAll(enable bool)
	SetTrackFixedSpeed(track int, percent int)

	GetMotoronsAttached() (tracks []bool)
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
