package keyin

import (
	"slotman/services/impl/provider"
	"slotman/services/type/keyin"
)

const (
	Service provider.Service = "serviceKeyin"
)

type Interface interface {
	GetName() (name provider.Service)

	Subscribe(subscriber keyin.Subscriber)
	Unsubscribe(subscriber keyin.Subscriber)
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
