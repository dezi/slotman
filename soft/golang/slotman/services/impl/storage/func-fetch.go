package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slotman/services/iface/identity"
	"slotman/services/type/storage"
	"slotman/utils/log"
	"sort"
	"strings"
	"time"
)

func (sv *Service) TryPlain(meta storage.Meta) (err error) {

	err = sv.fetchMeta(plain, meta, false)
	return
}

func (sv *Service) TryPlainFile(meta storage.File) (err error) {

	err = sv.fetchMeta(plain, meta, false)
	if err != nil {
		return
	}

	err = sv.fetchFile(plain, meta, false)
	return
}

func (sv *Service) TryPlainPath(meta storage.File) (storagePath string, err error) {

	storagePath, err = sv.fetchPath(plain, meta, false)
	if err != nil {
		return
	}

	err = sv.fetchMeta(plain, meta, false)
	return
}

func (sv *Service) FetchPlain(meta storage.Meta) (err error) {
	err = sv.fetchMeta(plain, meta, true)
	return
}

func (sv *Service) FetchPlainFile(meta storage.File) (err error) {

	err = sv.fetchMeta(plain, meta, true)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.fetchFile(plain, meta, true)
	return
}

func (sv *Service) FetchLatestPlain(meta storage.Meta) (err error) {

	ts, err := sv.retrieveLatest(plain, meta)
	if err != nil {
		if !errors.Is(err, errNoFilesFound) {
			log.Cerror(err)
		}
		return
	}

	meta.SetTime(&ts)

	err = sv.fetchMeta(plain, meta, true)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}

func (sv *Service) FetchLatestPlainFile(meta storage.File) (err error) {

	ts, err := sv.retrieveLatest(plain, meta)
	if err != nil {
		log.Cerror(err)
		return
	}

	meta.SetTime(&ts)

	err = sv.fetchMeta(plain, meta, true)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.fetchFile(plain, meta, true)
	return
}

func (sv *Service) fetchMeta(mode Mode, meta storage.Meta, logErr bool) (err error) {

	jsonFext := ".json"
	jsonBytes, err := sv.fetch(mode, meta, jsonFext, logErr)
	if err != nil {
		if logErr {
			log.Cerror(err)
		}
		return
	}

	err = json.Unmarshal(jsonBytes, meta)
	if err != nil {
		if logErr {
			log.Cerror(err)
		}
		return
	}

	return
}

func (sv *Service) fetchFile(mode Mode, meta storage.File, logErr bool) (err error) {

	dataFext := meta.GetFext()
	if !strings.HasPrefix(dataFext, ".") {
		dataFext = "." + dataFext
	}

	dataBytes, err := sv.fetch(mode, meta, dataFext, logErr)
	if err != nil {
		if logErr {
			log.Cerror(err)
		}
		return
	}

	meta.SetData(dataBytes)
	return
}

func (sv *Service) fetchPath(mode Mode, meta storage.File, logErr bool) (storagePath string, err error) {

	dataFext := meta.GetFext()
	storagePath, err = sv.fetchStoragePath(mode, meta, dataFext, logErr)
	if err != nil {
		if logErr {
			log.Cerror(err)
		}
		return
	}

	return
}

func (sv *Service) fetchStoragePath(mode Mode, meta storage.Meta, fext string, logErr bool) (storagePath string, err error) {

	caller := sv.getCallerPackage(9)
	if caller == "unknown" {
		err = errors.New("cannot resolve call stack")
		log.Cerror(err)
		return
	}

	storagePath = identity.GetStoragePath() +
		"/" + string(mode) +
		"/" + caller

	sub := meta.GetSub()

	if sub != "" {
		storagePath += "/" + sub
	}

	ts := meta.GetTime()
	day := meta.GetDay()

	if ts != nil && day {
		dayStr := ts.Format(datePartFormat)
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

	if ts != nil {
		if day {
			timeStr := ts.Format(timePartFormat)
			fileName += timeStr
		} else {
			dateStr := ts.Format(dateTimeFormat)
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

	_, err = os.Stat(storagePath)
	if err != nil {
		if logErr {
			log.Cerror(err)
		}
		return
	}

	return
}

func (sv *Service) fetch(mode Mode, meta storage.Meta, fext string, logErr bool) (data []byte, err error) {

	caller := sv.getCallerPackage(9)
	if caller == "unknown" {
		err = errors.New("cannot resolve call stack")
		log.Cerror(err)
		return
	}

	storagePath := identity.GetStoragePath() +
		"/" + string(mode) +
		"/" + caller

	ts := meta.GetTime()
	day := meta.GetDay()
	sub := meta.GetSub()

	if sub != "" {
		storagePath += "/" + sub
	}

	if ts != nil && day {
		dayStr := ts.Format(datePartFormat)
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

	if ts != nil {
		if day {
			timeStr := ts.Format(timePartFormat)
			fileName += timeStr
		} else {
			dateStr := ts.Format(dateTimeFormat)
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

	data, err = os.ReadFile(storagePath)
	if err != nil {
		if logErr {
			log.Cerror(err)
		}
		return
	}

	return
}

func (sv *Service) retrieveLatest(mode Mode, meta storage.Meta) (ts time.Time, err error) {

	caller := sv.getCallerPackage(7)
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

	day := meta.GetDay()

	if day {

		var datePart string
		datePart, err = sv.retrieveLatestDate(storagePath)
		if err != nil {
			log.Cerror(err)
			return
		}

		storagePath = filepath.Join(storagePath, datePart)

		var timePart string
		timePart, err = sv.retrieveLatestTime(storagePath, meta.GetTag())
		if err != nil {
			log.Cerror(err)
			return
		}

		dateTime := datePart + "-" + timePart

		ts, err = time.Parse(dateTimeFormat, dateTime)
		if err != nil {
			log.Cerror(err)
			return
		}

		return
	}

	var dateTime string
	dateTime, err = sv.retrieveLatestDateTime(storagePath, meta.GetTag())
	if err != nil {
		return
	}

	ts, err = time.Parse(dateTimeFormat, dateTime)
	if err == nil {
		return
	}

	ts, err = time.Parse(datePartFormat, dateTime)
	if err == nil {
		return
	}

	return
}

func (sv *Service) retrieveLatestDate(storagePath string) (datePart string, err error) {

	fis, err := os.ReadDir(storagePath)

	sort.Slice(fis, func(i, j int) bool {
		return fis[i].Name() > fis[j].Name()
	})

	for len(fis) > 0 && strings.HasPrefix(fis[0].Name(), ".") {
		fis = fis[1:]
	}

	if len(fis) == 0 {
		err = errNoFilesFound
		return
	}

	name := fis[0].Name()
	dfmt := datePartFormat

	if len(name) != len(dfmt) {
		err = errors.New("path not in day format")
		log.Cerror(err)
		return
	}

	name = name[:len(dfmt)]

	_, err = time.Parse(dfmt, name)
	if err != nil {
		return
	}

	datePart = name
	return
}

func (sv *Service) retrieveLatestTime(storagePath string, tag string) (timePart string, err error) {

	fis, err := os.ReadDir(storagePath)

	sort.Slice(fis, func(i, j int) bool {
		return fis[i].Name() > fis[j].Name()
	})

	for len(fis) > 0 && strings.HasPrefix(fis[0].Name(), ".") {
		fis = fis[1:]
	}

	if tag != "" {
		suffix := tag + ".json"
		for len(fis) > 0 && !strings.HasSuffix(fis[0].Name(), suffix) {
			fis = fis[1:]
		}
	}

	if len(fis) == 0 {
		err = errNoFilesFound
		return
	}

	name := fis[0].Name()
	tfmt := timePartFormat

	if len(name) < len(tfmt) {
		err = errors.New("file not in time format")
		log.Cerror(err)
		return
	}

	name = name[:len(tfmt)]

	_, err = time.Parse(tfmt, name)
	if err != nil {
		return
	}

	timePart = name
	return
}

func (sv *Service) retrieveLatestDateTime(storagePath string, tag string) (dateTime string, err error) {

	fis, err := os.ReadDir(storagePath)

	sort.Slice(fis, func(i, j int) bool {
		return fis[i].Name() > fis[j].Name()
	})

	for len(fis) > 0 && strings.HasPrefix(fis[0].Name(), ".") {
		fis = fis[1:]
	}

	if tag != "" {
		suffix := tag + ".json"
		for len(fis) > 0 && !strings.HasSuffix(fis[0].Name(), suffix) {
			fis = fis[1:]
		}
	}

	if len(fis) == 0 {
		err = errNoFilesFound
		return
	}

	name := fis[0].Name()

	if len(name) >= len(dateTimeFormat) {
		test := name[:len(dateTimeFormat)]
		_, err = time.Parse(dateTimeFormat, test)
		if err == nil {
			dateTime = test
			return
		}
	}

	if len(name) >= len(datePartFormat) {
		test := name[:len(datePartFormat)]
		_, err = time.Parse(datePartFormat, test)
		if err == nil {
			dateTime = test
			return
		}
	}

	err = errors.New("date latest error")
	return
}
