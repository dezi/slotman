package sgp40

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func NewSGP40(devicePath string) (se *SGP40) {
	se = &SGP40{
		Vendor:      "SENSIRION",
		Model:       "SGP40 co2 sensor",
		DevicePath:  devicePath,
		humidity:    50,
		temperature: 20,
	}
	return
}

func (se *SGP40) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *SGP40) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeAirQuality}
	return
}

func (se *SGP40) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *SGP40) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *SGP40) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *SGP40) Open() (err error) {

	shaData := fmt.Sprintf("%s|%s|%s|%s", things.ThingSystemUuid, se.Model, se.Vendor, se.DevicePath)
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

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *SGP40) Close() (err error) {

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

func (se *SGP40) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	go se.readLoop()

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *SGP40) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
