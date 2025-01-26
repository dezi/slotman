package aht20

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func NewAHT20(devicePath string) (se *AHT20) {
	se = &AHT20{
		Vendor:     "ASAIR",
		Model:      "AHT20 humidity and pressure sensor",
		DevicePath: devicePath,
	}
	return
}

func (se *AHT20) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *AHT20) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeHumidity, things.ThingTypeTemperature}
	return
}

func (se *AHT20) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *AHT20) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *AHT20) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *AHT20) Open() (err error) {

	shaData := fmt.Sprintf("%s|%s|%s|%s", simple.ZeroUuidHex(), se.Model, se.Vendor, se.DevicePath)
	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	parts := strings.Split(se.DevicePath, ":")
	devicePath := parts[0]

	i2cDev := i2c.NewDevice(devicePath, ThingI2CAddress)

	err = i2cDev.Open()
	if err != nil {
		return
	}

	se.i2cDev = i2cDev
	se.IsOpen = true

	se.SetThreshold(0.3)

	_ = se.Reset()

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *AHT20) Close() (err error) {

	if !se.IsOpen {
		return err
	}

	if se.IsStarted {
		_ = se.Stop()
	}

	se.IsOpen = false

	err = se.i2cDev.Close()
	log.Cerror(err)

	if se.handler != nil {
		se.handler.OnThingClosed(se)
	}

	return
}

func (se *AHT20) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	go se.readLoop()

	return
}

func (se *AHT20) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
