package pilots

import "slotman/services/impl/provider"

const (
	Provider provider.Provider = "servicePilots"
)

type Interface interface {
	GetName() (name provider.Provider)
}

func GetInstance() (iface Interface, err error) {

	baseProvider, err := provider.GetProvider(Provider)
	if err != nil {
		return
	}

	iface = baseProvider.(Interface)
	if iface == nil {
		err = provider.ErrNotFound(Provider)
		return
	}

	return
}
