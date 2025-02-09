package speedi

import (
	"slotman/services/impl/provider"
)

const (
	Service provider.Service = "serviceSpeedi"
)

type Interface interface {
	GetName() (name provider.Service)

	GetSpeedControlsAttached() (tracks []bool)
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
