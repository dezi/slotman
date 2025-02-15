package identity

import (
	"slotman/services/impl/provider"
	"slotman/utils/simple"
)

const (
	Service provider.Service = "serviceDisplay"
)

type Interface interface {
	GetName() (name provider.Service)

	GetBoxTag() (boxTag string)
	GetBoxIdentity() (boxIdentity simple.UUIDHex)
	GetStoragePath() (storagePath string)
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

//
// Interface shortcuts.
//

func GetBoxTag() (boxTag string) {
	id, _ := GetInstance()
	if id != nil {
		boxTag = id.GetBoxTag()
	}
	return
}

func GetBoxIdentity() (boxIdentity simple.UUIDHex) {
	id, _ := GetInstance()
	if id != nil {
		boxIdentity = id.GetBoxIdentity()
	}
	return
}

func GetStoragePath() (storagePath string) {
	id, _ := GetInstance()
	if id != nil {
		storagePath = id.GetStoragePath()
	}
	return
}
