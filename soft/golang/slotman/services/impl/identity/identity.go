package identity

import (
	"slotman/services/iface/identity"
	"slotman/services/impl/provider"
	"slotman/utils/log"
	"slotman/utils/simple"
)

type Service struct {
	globalBoxTag      string
	globalBoxIdentity simple.UUIDHex
	globalStoragePath string
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.globalBoxTag = "tch-iot"

	err = singleTon.setUpStoragePath()
	if err != nil {
		log.Cerror(err)
		return
	}

	err = singleTon.setUpBoxIdentity()
	if err != nil {
		log.Cerror(err)
		return
	}

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	log.Printf("Stopped service.")

	provider.UnsetProvider(singleTon)

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return identity.Service
}
