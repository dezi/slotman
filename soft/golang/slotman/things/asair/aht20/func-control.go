package aht20

import (
	"errors"
	"slotman/utils/log"
	"time"
)

func (se *AHT20) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *AHT20) SetThreshold(threshold float64) {
	se.threshold = threshold
	return
}

func (se *AHT20) Init() (err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	err = se.i2cDev.WriteRegBytes(byte(RegisterInit), []byte{0x08, 0x00})
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 50)
	return
}

func (se *AHT20) Reset() (err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	err = se.i2cDev.WriteReg(byte(RegisterReset))
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 50)
	return
}

func (se *AHT20) ReadMeasurement() (humidity, celsius float64, err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	err = se.i2cDev.WriteRegBytes(byte(RegisterMeasure), []byte{0x33, 0x00})
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 50)

	status := make([]byte, 7)
	xfer, err := se.i2cDev.ReadBytes(status)
	if err != nil {
		return
	}

	if xfer != 7 {
		err = errors.New("no measurement available")
		return
	}

	humidity = float64(int(status[1])<<12 | int(status[2])<<4 | int(status[3]&0xf0)>>4)
	humidity = humidity / 1048576 * 100

	celsius = float64(int(status[3]&0x0f)<<16 | int(status[4])<<8 | int(status[5]))
	celsius = celsius/1048576*200 - 50
	return
}
