package storage

import (
	"os"
	"slotman/services/impl/provider"
	"slotman/services/type/storage"
)

const (
	Service provider.Service = "serviceStorage"
)

type Interface interface {
	GetName() (name provider.Service)

	GetPathPlain() (storagePath string, err error)

	GetPlainStoragePathForPackage() (storagePath string)

	TryPlain(meta storage.Meta) (err error)
	TryPlainFile(meta storage.File) (err error)
	TryPlainPath(meta storage.File) (storagePath string, err error)

	FetchPlain(meta storage.Meta) (err error)
	FetchPlainFile(meta storage.File) (err error)
	FetchLatestPlain(meta storage.Meta) (err error)

	StorePlain(meta storage.Meta) (err error)
	StorePlainFile(meta storage.File) (err error)

	ListPlain(tag string) (fileInfos []os.FileInfo, err error)
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
