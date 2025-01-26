package bmp280

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func NewBMP280(devicePath string) (se *BMP280) {
	se = &BMP280{
		Vendor:     "BOSCH",
		Model:      "BMP280 digital pressure sensor",
		DevicePath: devicePath,
	}
	return
}

func (se *BMP280) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *BMP280) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypePressure, things.ThingTypeTemperature}
	return
}

func (se *BMP280) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *BMP280) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *BMP280) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *BMP280) Open() (err error) {

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

	_ = se.ResetSensor()
	_ = se.SetPowerMode(PowerModeNormal, PowerInterval1000ms)
	_ = se.SetMeasureMode(Oversampling16, Oversampling2)
	_ = se.readCompensation()

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *BMP280) Close() (err error) {

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

func (se *BMP280) Start() (err error) {

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

func (se *BMP280) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
