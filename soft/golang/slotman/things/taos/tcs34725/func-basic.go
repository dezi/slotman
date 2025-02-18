package tcs34725

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/services/iface/identity"
	"slotman/things"
	"slotman/utils/simple"
	"strings"
)

func NewTCS34275(devicePath string) (se *TCS34725) {
	se = &TCS34725{
		Vendor:     "TAOS",
		Model:      "TCS34725 RGB color sensor",
		DevicePath: devicePath,
	}
	return
}

func (se *TCS34725) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *TCS34725) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeRGBColor, things.ThingTypeIllumination}
	return
}

func (se *TCS34725) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *TCS34725) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *TCS34725) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *TCS34725) Open() (err error) {

	if se.IsOpen {
		return
	}

	shaData := fmt.Sprintf("%s|%s|%s|%s", identity.GetBoxIdentity(), se.Model, se.Vendor, se.DevicePath)
	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	parts := strings.Split(se.DevicePath, ":")
	devicePath := parts[0]

	i2cDev := i2c.NewDevice(devicePath, ThingAddress)

	err = i2cDev.Open()
	if err != nil {
		return
	}

	se.i2cDev = i2cDev
	se.IsOpen = true

	se.SetThreshold(3)

	err = se.InitThing(Gain1x, IntegrationTime101ms)
	if err != nil {
		_ = se.i2cDev.Close()
		se.i2cDev = nil
		se.IsOpen = false
		return
	}

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *TCS34725) Close() (err error) {

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

func (se *TCS34725) Start() (err error) {

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

func (se *TCS34725) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
