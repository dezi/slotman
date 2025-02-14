package server

import (
	"slotman/services/impl/provider"
	"slotman/services/type/server"
	"slotman/utils/simple"
)

const (
	Service provider.Service = "serviceServer"
)

type Interface interface {
	GetName() (name provider.Service)

	Subscribe(what string, handler server.Subscriber)
	Unsubscribe(what string)

	Transmit(appId simple.UUIDHex, resBytes []byte) (err error)
	Broadcast(appId simple.UUIDHex, resBytes []byte) (err error)
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
