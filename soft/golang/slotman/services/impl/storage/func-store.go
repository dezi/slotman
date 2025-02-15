package storage

import (
	"errors"
	"os"
	"slotman/services/iface/identity"
	"slotman/services/type/storage"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func (sv *Service) StorePlain(meta storage.Meta) (err error) {
	err = sv.storeMeta(plain, meta)
	return
}

func (sv *Service) StorePlainFile(meta storage.File) (err error) {
	err = sv.storeMeta(plain, meta)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.storeFile(plain, meta)
	return
}

func (sv *Service) storeMeta(mode Mode, meta storage.Meta) (err error) {

	jsonFext := ".json"
	jsonBytes, err := simple.MarshalJsonClean(meta)

	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.store(mode, meta, jsonFext, jsonBytes)
	return
}

func (sv *Service) storeFile(mode Mode, meta storage.File) (err error) {

	dataFext := meta.GetFext()
	if !strings.HasPrefix(dataFext, ".") {
		dataFext = "." + dataFext
	}

	dataBytes := meta.GetData()

	err = sv.store(mode, meta, dataFext, dataBytes)
	return
}

func (sv *Service) store(mode Mode, meta storage.Meta, fext string, data []byte) (err error) {

	caller := sv.getCallerPackage(9)
	if caller == "unknown" {
		err = errors.New("cannot resolve call stack")
		log.Cerror(err)
		return
	}

	storagePath := identity.GetStoragePath() +
		"/" + string(mode) +
		"/" + caller

	sub := meta.GetSub()

	if sub != "" {
		storagePath += "/" + sub
	}

	time := meta.GetTime()
	day := meta.GetDay()

	if time != nil && day {
		dayStr := time.Format(datePartFormat)
		storagePath += "/" + dayStr
	}

	_, statErr := os.Stat(storagePath)
	if statErr != nil {
		err = os.MkdirAll(storagePath, 0755)
		if err != nil {
			log.Cerror(err)
			return
		}
	}

	uuid := meta.GetUuid()
	tag := meta.GetTag()

	fileName := ""

	if time != nil {
		if day {
			timeStr := time.Format(timePartFormat)
			fileName += timeStr
		} else {
			dateStr := time.Format(dateTimeFormat)
			fileName += dateStr
		}
	}

	if uuid != nil {
		if len(fileName) > 0 {
			fileName += "."
		}
		fileName += string(*uuid)
	}

	if tag != "" {
		if len(fileName) > 0 {
			fileName += "."
		}
		fileName += tag
	}

	fileName += fext

	storagePath += "/" + fileName

	log.Debugf("storagePath=%s", storagePath)

	err = os.WriteFile(storagePath, data, 0644)

	return
}
