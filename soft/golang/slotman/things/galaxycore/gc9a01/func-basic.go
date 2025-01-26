package gc9a01

import (
	"errors"
	"fmt"
	"slotman/drivers/impl/gpio"
	"slotman/drivers/impl/spi"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func NewGC9A01(devicePath string, dcPinNo byte) (rc *GC9A01) {
	rc = &GC9A01{
		Vendor:     "GalaxyCore",
		Model:      "GC9A01 color display",
		DevicePath: devicePath,
		dcPinNo:    dcPinNo,
	}
	return
}

func (se *GC9A01) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *GC9A01) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeColorDisplay}
	return
}

func (se *GC9A01) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *GC9A01) GetThingAddress() (address int) {
	return
}

func (se *GC9A01) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *GC9A01) Open() (err error) {

	shaData := fmt.Sprintf("%s|%s|%s|%s", simple.ZeroUuidHex(), se.Model, se.Vendor, se.DevicePath)
	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	ok, err := gpio.HasGpio()
	if err != nil {
		return
	}

	if !ok {
		err = errors.New("no gpio available")
		return
	}

	se.dcPin = gpio.NewPin(25)

	err = se.dcPin.Open()
	if err != nil {
		return
	}

	err = se.dcPin.SetOutput()
	if err != nil {
		return
	}

	spiDev := spi.NewDevice(se.DevicePath)

	err = spiDev.Open()
	if err != nil {
		return
	}

	_ = spiDev.SetMode(0)
	_ = spiDev.SetBitsPerWord(8)
	_ = spiDev.SetSpeed(40000000)

	se.spiDev = spiDev

	err = se.Initialize()
	if err != nil {
		_ = se.spiDev.Close()
		se.spiDev = nil
		return
	}

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *GC9A01) Close() (err error) {

	if !se.IsOpen {
		return err
	}

	if se.IsStarted {
		_ = se.Stop()
	}

	se.IsOpen = false

	err = se.dcPin.Close()
	log.Cerror(err)

	err = se.spiDev.Close()
	log.Cerror(err)

	if se.handler != nil {
		se.handler.OnThingClosed(se)
	}

	return
}

func (se *GC9A01) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *GC9A01) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
