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
	_ = se.spi.SetSpeed(80000000)

	se.dcPin, err = gpio.GetPin(25)
	if err != nil {
		return
	}

	se.dcPin.SetOutput()

	return
}

func (se *GC9A01) SetFrame(frame Frame) (err error) {

	var data [4]byte

	data[0] = byte(frame.X0 >> 8)
	data[1] = byte(frame.X0)
	data[2] = byte(frame.X1 >> 8)
	data[3] = byte(frame.X1)

	_ = se.writeCommand(COL_ADDR_SET)
	_ = se.writeBytes(data[:])

	data[0] = byte(frame.Y0 >> 8)
	data[1] = byte(frame.Y0)
	data[2] = byte(frame.Y1 >> 8)
	data[3] = byte(frame.Y1)

	_ = se.writeCommand(ROW_ADDR_SET)
	_ = se.writeBytes(data[:])

	return
}
