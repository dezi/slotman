package pilots

import (
	"image"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
)

const (
	Service provider.Service = "servicePilots"
)

type Interface interface {
	GetName() (name provider.Service)

	GetAllPilots() (pilots []*slotman.Pilot)

	UpdatePilot(pilot *slotman.Pilot)
	GetScaledPilotPic(pilot *slotman.Pilot, size int) (img *image.RGBA, err error)
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
