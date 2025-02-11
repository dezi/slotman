package tacho

import (
	"slotman/services/impl/provider"
)

const (
	Service provider.Service = "serviceTacho"
)

type Interface interface {
	GetName() (name provider.Service)

	GetTachosAttached() (tracks []bool)

	OnRaceStarted()
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
