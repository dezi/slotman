package race

import (
	"slotman/services/impl/provider"
)

const (
	Service provider.Service = "serviceRace"
)

type Interface interface {
	GetName() (name provider.Service)

	OnAmpelClickShort()
	OnAmpelClickLong()
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
