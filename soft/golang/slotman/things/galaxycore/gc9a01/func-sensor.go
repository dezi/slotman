package gc9a01

import (
	"slotman/drivers/gpio"
	"slotman/drivers/spi"
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

func (se *GC9A01) OpenSensor() (err error) {

	//shaData := fmt.Sprintf("%s|%s|%s|%s", identity.GetBoxIdentity(), se.Model, se.Vendor, se.DevicePath)
	//se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	se.dcPin, err = gpio.GetPin(25)
	if err != nil {
		return
	}

	se.dcPin.SetOutput()

	spiDev := spi.NewDevice(se.DevicePath)

	err = spiDev.Open()
	if err != nil {
		return
	}

	_ = spiDev.SetMode(0)
	_ = spiDev.SetBitsPerWord(8)
	_ = spiDev.SetSpeed(80000000)

	se.spi = spiDev

	err = se.Initialize()
	if err != nil {
		_ = se.spi.Close()
		se.spi = nil
		return
	}

	return
}
