package identity

import (
	"errors"
	"os"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) setUpStoragePath() (err error) {

	storagePath, err := simple.ResolvePath("~/" + sv.globalBoxTag)
	if err != nil {
		log.Cerror(err)
		return
	}

	_, err = os.Stat(storagePath)
	if err != nil {
		err = os.Mkdir(storagePath, 0755)
		if err != nil {
			log.Cerror(err)
			return
		}
	}

	sv.globalStoragePath = storagePath

	log.Printf("Global storagePath=%s", sv.globalStoragePath)

	return
}

func (sv *Service) setUpBoxIdentity() (err error) {

	storageIdentity, storageErr := searchIdentity(sv.globalStoragePath)

	if storageErr == nil {
		sv.globalBoxIdentity = storageIdentity
		log.Printf("Global boxIdentity=%s", sv.globalBoxIdentity)
		return
	}

	//
	// Create a new identity.
	//

	newBoxIdentity := simple.UuidToHexString(simple.NewUuid())

	storageIdentityPath := sv.globalStoragePath + "/" + string(newBoxIdentity)
	err = os.Mkdir(storageIdentityPath, 0755)
	if err != nil {
		log.Cerror(err)
		return
	}

	sv.globalBoxIdentity = newBoxIdentity
	log.Printf("Created boxIdentity=%s", sv.globalBoxIdentity)

	return
}

func searchIdentity(path string) (identity simple.UUIDHex, err error) {

	if path == "" {
		err = errors.New("no path")
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Cerror(err)
		return
	}

	for _, file := range files {

		uuid, parseErr := simple.UuidFromHexString(simple.UUIDHex(file.Name()))
		if parseErr != nil {
			continue
		}

		identity = simple.UuidToHexString(uuid)
		return
	}

	err = errors.New("no identity found")
	return
}
