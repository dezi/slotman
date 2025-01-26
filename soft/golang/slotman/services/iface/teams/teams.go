package teams

import (
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
)

const (
	Service provider.Service = "serviceTeams"
)

type Interface interface {
	GetName() (name provider.Service)

	GetAllTeams() (teams []*slotman.Team)
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
