package ads1115

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"strings"
)

func NewADS1115(devicePath string) (se *ADS1115) {

	se = &ADS1115{
		Vendor:     "TI",
		Model:      "ADS1115 ADC",
		DevicePath: devicePath,
	}

	_ = se.SetGain(0, Gain2)
	_ = se.SetGain(1, Gain2)
	_ = se.SetGain(2, Gain2)
	_ = se.SetGain(3, Gain2)

	_ = se.SetRate(0, Rate128ps)
	_ = se.SetRate(1, Rate128ps)
	_ = se.SetRate(2, Rate128ps)
	_ = se.SetRate(3, Rate128ps)
	return
}

func (se *ADS1115) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *ADS1115) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeADConverter}
	return
}

func (se *ADS1115) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *ADS1115) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *ADS1115) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *ADS1115) Open() (err error) {

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

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *ADS1115) Close() (err error) {

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

func (se *ADS1115) Start() (err error) {

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

func (se *ADS1115) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
