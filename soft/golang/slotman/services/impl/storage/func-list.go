package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slotman/services/iface/identity"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func (sv *Service) ListPlain(tag string) (fileInfos []os.FileInfo, err error) {
	return sv.list(plain, tag)
}

func (sv *Service) list(mode Mode, tag string) (fileInfos []os.FileInfo, err error) {

	if tag != "" {
		tag = "." + tag + "."
	}

	storagePath := identity.GetStoragePath() +
		"/" + string(mode) +
		"/" + sv.getCallerPackage(7)

	des, dirErr := os.ReadDir(storagePath)
	if dirErr != nil {
		return
	}

	for _, de := range des {

		if de.Name() == ".DS_Store" {
			continue
		}

		if tag != "" && !strings.Contains(de.Name(), tag) {
			continue
		}

		var fi os.FileInfo
		fi, err = de.Info()
		if err != nil {
			return
		}

		fileInfos = append(fileInfos, fi)
	}

	return
}

func retrieveFilename(jsonBytes []byte) (fileName string, err error) {

	var jsonObj interface{}
	err = json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		log.Cerror(err)
		return
	}

	topMap := jsonObj.(map[string]interface{})

	uuid, ok := topMap["Uuid"]
	if !ok {
		uuid, ok = topMap["uuid"]
		if !ok {
			err = errors.New("no UUID in data")
			return
		}
	}

	uuidHex := ""

	switch uuid.(type) {
	case string:
		uuidHex = uuid.(string)
	}

	_, err = simple.UuidFromHexString(simple.UUIDHex(uuidHex))
	if err != nil {
		err = errors.New(fmt.Sprintf("bad UUID format <%s>", uuidHex))
		log.Cerror(err)
		return
	}

	fileName = uuidHex
	return
}
