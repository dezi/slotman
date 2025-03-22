package ld6001a

import (
	"fmt"
	"slotman/drivers/impl/uart"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
	"time"
)

func NewLD6001a(devicePath string, baudRate int) (se *LD6001a) {
	se = &LD6001a{
		Vendor:     "HI-Link",
		Model:      "LD6001a 24GHz Human Tracking",
		DevicePath: devicePath,
		BaudRate:   baudRate,
	}
	return
}

func (se *LD6001a) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *LD6001a) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeHumanTrack}
	return
}

func (se *LD6001a) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *LD6001a) GetThingAddress() (address int) {
	return
}

func (se *LD6001a) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *LD6001a) Open() (err error) {

	shaData := fmt.Sprintf("%s|%s|%s|%s", things.ThingSystemUuid, se.Model, se.Vendor, se.DevicePath)
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
	se.results = make(chan string, 10)
	se.IsOpen = true

	se.loopGroup.Add(1)
	go se.readLoop()

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *LD6001a) Close() (err error) {

	if !se.IsOpen {
		return err
	}

	if se.IsStarted {
		_ = se.Stop()
	}

	se.IsOpen = false

	se.loopGroup.Wait()

	err = se.uart.Close()
	log.Cerror(err)

	if se.handler != nil {
		se.handler.OnThingClosed(se)
	}

	return
}

func (se *LD6001a) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *LD6001a) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
