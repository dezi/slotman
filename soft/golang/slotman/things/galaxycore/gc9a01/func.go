package gc9a01

import (
	"slotman/drivers/gpio"

	"slotman/drivers/spi"
	"slotman/utils/log"
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

	spiDev := spi.NewDevice(se.DevicePath)

	err = spiDev.Open()
	if err != nil {
		log.Cerror(err)
		return
	}

	se.spi = spiDev

	_ = se.spi.SetMode(0)
	_ = se.spi.SetBitsPerWord(8)
	_ = se.spi.SetSpeed(1000000)

	se.dcPin, err = gpio.GetPin(25)
	if err != nil {
		return
	}

	se.dcPin.SetOutput()

	return
}

func (se *GC9A01) WriteCommand(cmd byte) (err error) {

	se.dcPin.SetLow()

	_, err = se.spi.Send([]byte{cmd})

	return
}

func (se *GC9A01) WriteData(data []byte) (err error) {

	se.dcPin.SetHigh()

	_, err = se.spi.Send(data)

	return
}
