package ld2461

import (
	"fmt"
	"slotman/drivers/impl/uart"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
	"time"
)

func NewLD2461(devicePath string, baudRate int) (se *LD2461) {
	se = &LD2461{
		Vendor:     "HI-Link",
		Model:      "LD2461 24GHz Human Tracking",
		DevicePath: devicePath,
		BaudRate:   baudRate,
	}
	return
}

func (se *LD2461) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *LD2461) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeHumanTrack}
	return
}

func (se *LD2461) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *LD2461) GetThingAddress() (address int) {
	return
}

func (se *LD2461) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *LD2461) Open() (err error) {

	shaData := fmt.Sprintf("%s|%s|%s|%s", simple.ZeroUuidHex(), se.Model, se.Vendor, se.DevicePath)
	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	uartPort := uart.NewDevice(se.DevicePath, se.BaudRate)
	err = uartPort.Open()
	if err != nil {
		return
	}

	err = uartPort.SetReadTimeout(time.Millisecond * 100)
	if err != nil {
		log.Cerror(err)
		return
	}

	se.uart = uartPort
	se.results = make(chan []byte, 10)
	se.IsOpen = true

	go se.readLoop()

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *LD2461) Close() (err error) {

	if !se.IsOpen {
		return err
	}

	if se.IsStarted {
		_ = se.Stop()
	}

	se.IsOpen = false

	err = se.uart.Close()
	log.Cerror(err)

	if se.handler != nil {
		se.handler.OnThingClosed(se)
	}

	return
}

func (se *LD2461) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *LD2461) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
