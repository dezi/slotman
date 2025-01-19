package gc9a01

import (
	"slotman/drivers/gpio"
	"time"

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
	_ = se.spi.SetSpeed(8000000)

	se.dcPin, err = gpio.GetPin(25)
	if err != nil {
		return
	}

	se.dcPin.SetOutput()

	//for {
	//	se.dcPin.SetLow()
	//	time.Sleep(time.Second * 5)
	//	se.dcPin.SetHigh()
	//	time.Sleep(time.Second * 2)
	//}

	return
}

func (se *GC9A01) WriteCommand(cmd byte) (err error) {

	se.dcPin.SetLow()

	_, err = se.spi.Send([]byte{cmd})

	return
}

func (se *GC9A01) WriteByte(data byte) (err error) {

	se.dcPin.SetHigh()

	_, err = se.spi.Send([]byte{data})

	return
}

func (se *GC9A01) WriteBytes(data []byte) (err error) {

	se.dcPin.SetHigh()

	_, err = se.spi.Send(data)

	return
}

func (se *GC9A01) WriteMem(data []byte) (err error) {
	err = se.WriteCommand(MEM_WR)
	err = se.WriteBytes(data)
	return
}

func (se *GC9A01) WriteMemCont(data []byte) (err error) {
	err = se.WriteCommand(MEM_WR_CONT)
	err = se.WriteBytes(data)
	return
}

func (se *GC9A01) Initialize() (err error) {

	_ = se.WriteCommand(0xEF)

	_ = se.WriteCommand(0xEB)
	_ = se.WriteByte(0x14)

	_ = se.WriteCommand(0xFE)
	_ = se.WriteCommand(0xEF)

	_ = se.WriteCommand(0xEB)
	_ = se.WriteByte(0x14)

	_ = se.WriteCommand(0x84)
	_ = se.WriteByte(0x40)

	_ = se.WriteCommand(0x85)
	_ = se.WriteByte(0xFF)

	_ = se.WriteCommand(0x86)
	_ = se.WriteByte(0xFF)

	_ = se.WriteCommand(0x87)
	_ = se.WriteByte(0xFF)

	_ = se.WriteCommand(0x88)
	_ = se.WriteByte(0x0A)

	_ = se.WriteCommand(0x89)
	_ = se.WriteByte(0x21)

	_ = se.WriteCommand(0x8A)
	_ = se.WriteByte(0x00)

	_ = se.WriteCommand(0x8B)
	_ = se.WriteByte(0x80)

	_ = se.WriteCommand(0x8C)
	_ = se.WriteByte(0x01)

	_ = se.WriteCommand(0x8D)
	_ = se.WriteByte(0x01)

	_ = se.WriteCommand(0x8E)
	_ = se.WriteByte(0xFF)

	_ = se.WriteCommand(0x8F)
	_ = se.WriteByte(0xFF)

	_ = se.WriteCommand(0xB6)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x00)

	_ = se.WriteCommand(0x36)
	_ = se.WriteByte(0x18)

	// #define COLOR_MODE          0x3A
	// #define COLOR_MODE__12_BIT  0x03
	// #define COLOR_MODE__16_BIT  0x05
	// #define COLOR_MODE__18_BIT  0x06

	_ = se.WriteCommand(COLOR_MODE)
	_ = se.WriteByte(COLOR_MODE_18_BIT)

	_ = se.WriteCommand(0x90)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x08)

	_ = se.WriteCommand(0xBD)
	_ = se.WriteByte(0x06)

	_ = se.WriteCommand(0xBC)
	_ = se.WriteByte(0x00)

	_ = se.WriteCommand(0xFF)
	_ = se.WriteByte(0x60)
	_ = se.WriteByte(0x01)
	_ = se.WriteByte(0x04)

	_ = se.WriteCommand(0xC3)
	_ = se.WriteByte(0x13)
	_ = se.WriteCommand(0xC4)
	_ = se.WriteByte(0x13)

	_ = se.WriteCommand(0xC9)
	_ = se.WriteByte(0x22)

	_ = se.WriteCommand(0xBE)
	_ = se.WriteByte(0x11)

	_ = se.WriteCommand(0xE1)
	_ = se.WriteByte(0x10)
	_ = se.WriteByte(0x0E)

	_ = se.WriteCommand(0xDF)
	_ = se.WriteByte(0x21)
	_ = se.WriteByte(0x0c)
	_ = se.WriteByte(0x02)

	_ = se.WriteCommand(0xF0)
	_ = se.WriteByte(0x45)
	_ = se.WriteByte(0x09)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x26)
	_ = se.WriteByte(0x2A)

	_ = se.WriteCommand(0xF1)
	_ = se.WriteByte(0x43)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x72)
	_ = se.WriteByte(0x36)
	_ = se.WriteByte(0x37)
	_ = se.WriteByte(0x6F)

	_ = se.WriteCommand(0xF2)
	_ = se.WriteByte(0x45)
	_ = se.WriteByte(0x09)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x26)
	_ = se.WriteByte(0x2A)

	_ = se.WriteCommand(0xF3)
	_ = se.WriteByte(0x43)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x72)
	_ = se.WriteByte(0x36)
	_ = se.WriteByte(0x37)
	_ = se.WriteByte(0x6F)

	_ = se.WriteCommand(0xED)
	_ = se.WriteByte(0x1B)
	_ = se.WriteByte(0x0B)

	_ = se.WriteCommand(0xAE)
	_ = se.WriteByte(0x77)

	_ = se.WriteCommand(0xCD)
	_ = se.WriteByte(0x63)

	_ = se.WriteCommand(0x70)
	_ = se.WriteByte(0x07)
	_ = se.WriteByte(0x07)
	_ = se.WriteByte(0x04)
	_ = se.WriteByte(0x0E)
	_ = se.WriteByte(0x0F)
	_ = se.WriteByte(0x09)
	_ = se.WriteByte(0x07)
	_ = se.WriteByte(0x08)
	_ = se.WriteByte(0x03)

	_ = se.WriteCommand(0xE8)
	_ = se.WriteByte(0x34)

	_ = se.WriteCommand(0x62)
	_ = se.WriteByte(0x18)
	_ = se.WriteByte(0x0D)
	_ = se.WriteByte(0x71)
	_ = se.WriteByte(0xED)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x18)
	_ = se.WriteByte(0x0F)
	_ = se.WriteByte(0x71)
	_ = se.WriteByte(0xEF)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x70)

	_ = se.WriteCommand(0x63)
	_ = se.WriteByte(0x18)
	_ = se.WriteByte(0x11)
	_ = se.WriteByte(0x71)
	_ = se.WriteByte(0xF1)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x18)
	_ = se.WriteByte(0x13)
	_ = se.WriteByte(0x71)
	_ = se.WriteByte(0xF3)
	_ = se.WriteByte(0x70)
	_ = se.WriteByte(0x70)

	_ = se.WriteCommand(0x64)
	_ = se.WriteByte(0x28)
	_ = se.WriteByte(0x29)
	_ = se.WriteByte(0xF1)
	_ = se.WriteByte(0x01)
	_ = se.WriteByte(0xF1)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x07)

	_ = se.WriteCommand(0x66)
	_ = se.WriteByte(0x3C)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0xCD)
	_ = se.WriteByte(0x67)
	_ = se.WriteByte(0x45)
	_ = se.WriteByte(0x45)
	_ = se.WriteByte(0x10)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x00)

	_ = se.WriteCommand(0x67)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x3C)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x01)
	_ = se.WriteByte(0x54)
	_ = se.WriteByte(0x10)
	_ = se.WriteByte(0x32)
	_ = se.WriteByte(0x98)

	_ = se.WriteCommand(0x74)
	_ = se.WriteByte(0x10)
	_ = se.WriteByte(0x85)
	_ = se.WriteByte(0x80)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x00)
	_ = se.WriteByte(0x4E)
	_ = se.WriteByte(0x00)

	_ = se.WriteCommand(0x98)
	_ = se.WriteByte(0x3e)
	_ = se.WriteByte(0x07)

	_ = se.WriteCommand(0x35)
	_ = se.WriteCommand(0x21)

	_ = se.WriteCommand(0x11)
	time.Sleep(time.Millisecond * 120)

	_ = se.WriteCommand(0x29)
	time.Sleep(time.Millisecond * 20)

	return
}

func (se *GC9A01) SetFrame(frame Frame) (err error) {

	var data [4]byte

	data[0] = byte(frame.X0 >> 8)
	data[1] = byte(frame.X0)
	data[2] = byte(frame.X1 >> 8)
	data[3] = byte(frame.X1)

	_ = se.WriteCommand(COL_ADDR_SET)
	_ = se.WriteBytes(data[:])

	data[0] = byte(frame.Y0 >> 8)
	data[1] = byte(frame.Y0)
	data[2] = byte(frame.Y1 >> 8)
	data[3] = byte(frame.Y1)

	_ = se.WriteCommand(ROW_ADDR_SET)
	_ = se.WriteBytes(data[:])

	return
}
