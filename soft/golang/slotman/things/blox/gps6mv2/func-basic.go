package gps6mv2

import (
	"fmt"
	"slotman/drivers/impl/uart"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
	"time"
)

func NewGPS6MV2(devicePath string, baudRate int) (se *GPS6MV2) {
	se = &GPS6MV2{
		Vendor:     "blox",
		Model:      "GPS6MV2 GPS module",
		DevicePath: devicePath,
		BaudRate:   baudRate,
	}
	return
}

func (se *GPS6MV2) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *GPS6MV2) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeGpsPosition}
	return
}

func (se *GPS6MV2) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *GPS6MV2) GetThingAddress() (address int) {
	return
}

func (se *GPS6MV2) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *GPS6MV2) Open() (err error) {

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

	go se.evalLoop()
	go se.readLoop()

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *GPS6MV2) Close() (err error) {

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

func (se *GPS6MV2) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *GPS6MV2) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
