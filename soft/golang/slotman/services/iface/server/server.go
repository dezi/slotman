package server

import (
	"slotman/services/impl/provider"
)

const (
	Service provider.Service = "serviceServer"
)

type Interface interface {
	GetName() (name provider.Service)

	Broadcast(resBytes []byte) (err error)
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
