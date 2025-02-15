package storage

import (
	"os"
	"slotman/services/iface/identity"
	"slotman/services/iface/storage"
	"slotman/services/impl/provider"
	"slotman/utils/log"
)

type Service struct {
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	err = setupStorageDirectories(identity.GetStoragePath())
	if err != nil {
		log.Cerror(err)
		return
	}

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return storage.Service
}

func setupStorageDirectories(path string) (err error) {

	if path == "" {
		return
	}

	err = setupStorageDirectory(path + "/" + string(plain))
	if err != nil {
		return
	}

	return
}

func setupStorageDirectory(path string) (err error) {

	_, statErr := os.Stat(path)
	if statErr == nil {
		return
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}
