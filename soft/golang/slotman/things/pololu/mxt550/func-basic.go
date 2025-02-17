package mxt550

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"strings"
)

func NewMXT550(devicePath string) (se *MXT550) {
	se = &MXT550{
		Vendor:     "Pololu",
		Model:      "MXT550",
		DevicePath: devicePath,
	}
	return
}

func (se *MXT550) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *MXT550) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeMotorDriver}
	return
}

func (se *MXT550) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *MXT550) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *MXT550) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *MXT550) Open() (err error) {

	if se.IsOpen {
		return
	}

	shaData := fmt.Sprintf("%s|%s|%s|%s", things.ThingSystemUuid, se.Model, se.Vendor, se.DevicePath)

	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	parts := strings.Split(se.DevicePath, ":")
	devicePath := parts[0]
	var address int
	_, _ = fmt.Sscanf(parts[1], "%02x", &address)

	i2cDev := i2c.NewDevice(devicePath, byte(address))

	err = i2cDev.Open()
	if err != nil {
		return
	}

	se.i2cDev = i2cDev
	se.IsOpen = true
	se.debug = false

	err = se.Reinitialize()
	if err != nil {
		_ = se.i2cDev.Close()
		se.IsOpen = false
		return
	}

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *MXT550) Close() (err error) {

	if !se.IsOpen {
		return
	}

	if se.IsStarted {
		_ = se.Stop()
	}

	se.IsOpen = false

	_ = se.i2cDev.Close()
	se.i2cDev = nil

	if se.handler != nil {
		se.handler.OnThingClosed(se)
	}

	return
}

func (se *MXT550) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *MXT550) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
