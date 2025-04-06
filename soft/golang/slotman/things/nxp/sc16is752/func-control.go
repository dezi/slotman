package sc16is752

import (
	"errors"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SC15IS752) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *SC15IS752) SetFifoEnable(channel byte, enable bool) (err error) {

	value, err := se.ReadRegister(RegFCR, channel)
	if err != nil {
		return
	}

	if enable {
		value |= 0x01
	} else {
		value &= 0xFE
	}

	err = se.WriteRegister(RegFCR, channel, value)
	return
}

func (se *SC15IS752) SetCrystalFreq(crystalFreq int) (err error) {
	se.crystalFreq = crystalFreq
	return
}

func (se *SC15IS752) SetBaudrate(channel byte, baudrate int) (err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	if se.baudrate[channel] == baudrate {
		return
	}

	se.baudrate[channel] = baudrate

	value, err := se.ReadRegister(RegMCR, channel)
	if err != nil {
		return
	}

	preScaler := 4
	if value&0x80 == 0 {
		preScaler = 1
	}

	divisor1 := se.crystalFreq / preScaler
	divisor2 := baudrate * 16
	if divisor2 > divisor1 {
		err = things.ErrUnsupportedBaudRate
		return
	}

	wk := float64(divisor1) / float64(divisor2)
	divisor := int(wk + 0.999)

	value, err = se.ReadRegister(RegLCR, channel)
	if err != nil {
		return
	}

	value |= 0x80

	err = se.WriteRegister(RegLCR, channel, value)
	if err != nil {
		return
	}

	err = se.WriteRegister(RegDLH, channel, byte(divisor>>8))
	if err != nil {
		return
	}

	err = se.WriteRegister(RegDLL, channel, byte(divisor))
	if err != nil {
		return
	}

	value &= 0x7F

	err = se.WriteRegister(RegLCR, channel, value)
	return
}

func (se *SC15IS752) SetLine(channel byte, dataBits, parity, stopBits byte) (err error) {

	if dataBits < 5 || dataBits > 8 {
		err = errors.New("invalid data bits")
		return
	}

	if parity > ParityForce0 {
		err = errors.New("invalid parity")
		return
	}

	if stopBits != 1 && stopBits != 2 {
		err = errors.New("invalid stop bits")
		return
	}

	value, err := se.ReadRegister(RegLCR, channel)
	if err != nil {
		return
	}

	value &= 0xC0

	//goland:noinspection GoDfaConstantCondition
	switch dataBits {
	case 5:
		value |= 0x00
	case 6:
		value |= 0x01
	case 7:
		value |= 0x02
	case 8:
		value |= 0x03
	}

	switch parity {
	case ParityNone:
		value |= 0x00
	case ParityOdd:
		value |= 0x08
	case ParityEven:
		value |= 0x18
	case ParityForce1:
		value |= 0x03
	case ParityForce0:
		value |= 0x00
	}

	//goland:noinspection GoDfaConstantCondition
	switch stopBits {
	case 1:
		value |= 0x00
	case 2:
		value |= 0x04
	}

	err = se.WriteRegister(RegLCR, channel, value)
	return
}

func (se *SC15IS752) SetPollInterval(channel byte, millis int) (err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.pollSleep[channel] = millis
	return
}

func (se *SC15IS752) SetReadTimeout(channel byte, millis int) (err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.readTimeout[channel] = millis
	return
}

func (se *SC15IS752) GetReadFifoAvail(channel byte) (avail int, err error) {

	value, err := se.ReadRegister(RegRxLvl, channel)
	if err != nil {
		return
	}

	avail = int(value)
	return
}

func (se *SC15IS752) GetWriteFifoAvail(channel byte) (avail int, err error) {

	value, err := se.ReadRegister(RegTxLvl, channel)
	if err != nil {
		return
	}

	avail = int(value)
	return
}

func (se *SC15IS752) WriteUartBytes(channel byte, data []byte) (xfer int, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	xfer, err = se.i2cDev.WriteUart(channel, 100, data)

	return
}

func (se *SC15IS752) WriteUartBytesLocal(channel byte, data []byte) (xfer int, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	//log.Printf("####### WriteUartBytes size=%d [ %02x ]", len(data), data)

	var avail int

	for xfer < len(data) {

		for {

			avail, err = se.GetWriteFifoAvail(channel)
			if err != nil {
				return
			}

			if avail > 0 {
				break
			}

			time.Sleep(time.Millisecond * 1)
		}

		if avail > len(data)-xfer {
			avail = len(data) - xfer
		}

		err = se.WriteRegisterBytes(RegTHR, channel, data[xfer:xfer+avail])
		if err != nil {
			log.Cerror(err)
			return
		}

		xfer += avail
	}

	//log.Printf("####### WriteUartBytes xfer=%d [ %02x ]", xfer, data)

	return
}

func (se *SC15IS752) ReadUartBytes(channel byte, size int) (xfer int, data []byte, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	data = make([]byte, size)

	xfer, err = se.i2cDev.ReadUart(channel, 25, data)
	data = data[:xfer]

	return
}

func (se *SC15IS752) ReadUartBytesLocal(channel byte, size int) (xfer int, data []byte, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	var avail int
	var timeOut = se.readTimeout[channel]
	var startTime = time.Now().UnixMilli()

	log.Printf("############# ReadUartBytes size=%d", size)

	for size > len(data) {

		for {

			avail, err = se.GetReadFifoAvail(channel)
			if err != nil {
				log.Cerror(err)
				return
			}

			if avail > 0 {
				log.Printf("######fifo avail=%d", avail)
				break
			}

			if timeOut > 0 && time.Now().UnixMilli()-startTime > int64(timeOut) {
				err = ErrReadTimeout
				log.Cerror(err)
				return
			}
		}

		if avail > size-len(data) {
			avail = size - len(data)
		}

		var read []byte
		read, _, err = se.ReadRegisterBytes(RegRHR, channel, avail)
		if err != nil {
			log.Printf("############### fuck here....")
			log.Cerror(err)
			return
		}

		data = append(data, read...)
		xfer = len(data)

		startTime = time.Now().UnixMilli()
	}

	log.Printf("############# ReadUartBytes size=%d xfer=%d done...", size, xfer)

	return
}

func (se *SC15IS752) ReadUartBytesNow(channel byte) (data []byte, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	var avail int

	avail, err = se.GetReadFifoAvail(channel)
	if err != nil {
		return
	}

	if avail == 0 {
		return
	}

	//var value byte
	//value, err = se.ReadRegister(RegRHR, channel)
	//if err != nil {
	//	return
	//}
	//
	//data = append(data, value)
	//
	//avail--

	data, _, err = se.ReadRegisterBytes(RegRHR, channel, avail)
	return
}

func (se *SC15IS752) Ping() (err error) {

	err = se.WriteRegister(RegSPR, ChannelA, 0x55)
	if err != nil {
		return
	}

	value, err := se.ReadRegister(RegSPR, ChannelA)
	if err != nil {
		return
	}

	if value != 0x55 {
		err = ErrInvalidPing
		return
	}

	err = se.WriteRegister(RegSPR, ChannelA, 0xaa)
	if err != nil {
		return
	}

	value, err = se.ReadRegister(RegSPR, ChannelA)
	if err != nil {
		return
	}

	if value != 0xaa {
		err = ErrInvalidPing
		return
	}

	return
}
