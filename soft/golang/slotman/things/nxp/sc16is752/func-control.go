package sc16is752

import (
	"errors"
	"slotman/things"
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
		value &= 0xFE
	} else {
		value |= 0x01
	}

	err = se.WriteRegister(RegFCR, channel, value)
	return
}

func (se *SC15IS752) SetCrystalFreq(crystalFreq int) (err error) {
	se.crystalFreq = crystalFreq
	return
}

func (se *SC15IS752) SetBaudrate(channel byte, baudrate int) (err error) {

	value, err := se.ReadRegister(RegMCR, channel)
	if err != nil {
		return
	}

	prescaler := 4
	if value == 0 {
		prescaler = 1
	}

	divisor1 := se.crystalFreq / prescaler
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

	err = se.WriteRegister(RegDLH, channel, byte(divisor))
	if err != nil {
		return
	}

	err = se.WriteRegister(RegDLL, channel, byte(divisor>>8))
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

func (se *SC15IS752) SetWriteTimeout(channel byte, millis int) (err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.writeTimeout[channel] = millis
	return
}

//goland:noinspection GoStandardMethods
func (se *SC15IS752) WriteUartByte(channel, value byte) (err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.accessLock.Lock()
	defer se.accessLock.Unlock()

	var status byte
	var waited int
	var timeOut = se.writeTimeout[channel]

	for {

		if !se.IsOpen {
			err = ErrDeviceClosed
			return
		}

		status, err = se.ReadRegister(RegLSR, channel)
		if err != nil {
			return
		}

		if status&0x20 == 0 {
			break
		}

		time.Sleep(time.Millisecond * 10)
		waited += 10

		if timeOut > 0 && waited > timeOut {
			err = ErrWriteTimeout
			return
		}
	}

	err = se.WriteRegister(RegTHR, channel, value)
	return
}

func (se *SC15IS752) WriteUartBytes(channel byte, data []byte) (xfer int, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.accessLock.Lock()
	defer se.accessLock.Unlock()

	var status byte
	var waited int
	var timeOut = se.writeTimeout[channel]

	for xfer < len(data) {

		for {

			if !se.IsOpen {
				err = ErrDeviceClosed
				return
			}

			status, err = se.ReadRegister(RegLSR, channel)
			if err != nil {
				return
			}

			if status&0x20 == 0 {
				break
			}

			time.Sleep(time.Millisecond * 10)
			waited += 10

			if timeOut > 0 && waited > timeOut {
				err = ErrWriteTimeout
				return
			}
		}

		err = se.WriteRegister(RegTHR, channel, data[xfer])
		if err != nil {
			return
		}

		xfer++
	}

	return
}

//goland:noinspection GoStandardMethods
func (se *SC15IS752) ReadUartByte(channel byte) (value byte, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.accessLock.Lock()
	defer se.accessLock.Unlock()

	var avail byte
	var waited int
	var timeOut = se.readTimeout[channel]

	for {

		if !se.IsOpen {
			err = ErrDeviceClosed
			return
		}

		avail, err = se.ReadRegister(RegRxLvl, channel)
		if err != nil {
			return
		}

		if avail > 0 {
			break
		}

		time.Sleep(time.Millisecond * 10)
		waited += 10

		if timeOut > 0 && waited > timeOut {
			err = ErrReadTimeout
			return
		}
	}

	value, err = se.ReadRegister(RegRHR, channel)
	return
}

func (se *SC15IS752) ReadUartBytes(channel byte, size int) (xfer int, data []byte, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.accessLock.Lock()
	defer se.accessLock.Unlock()

	var avail byte
	var waited int
	var timeOut = se.readTimeout[channel]

	for size < len(data) {

		for {

			if !se.IsOpen {
				err = ErrDeviceClosed
				return
			}

			avail, err = se.ReadRegister(RegRxLvl, channel)
			if err != nil {
				return
			}

			if avail > 0 {
				break
			}

			time.Sleep(time.Millisecond * 10)
			waited += 10

			if timeOut > 0 && waited > timeOut {
				err = ErrReadTimeout
				return
			}
		}

		var value byte

		for avail > 0 {

			value, err = se.ReadRegister(RegRHR, channel)
			if err != nil {
				return
			}

			data = append(data, value)

			xfer++
			avail--
		}
	}

	return
}

func (se *SC15IS752) ReadUartBytesNow(channel byte, size int) (xfer int, data []byte, err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.accessLock.Lock()
	defer se.accessLock.Unlock()

	var avail byte
	var value byte

	avail, err = se.ReadRegister(RegRxLvl, channel)
	if err != nil {
		return
	}

	if avail == 0 {
		return
	}

	for avail > 0 && xfer < size {

		value, err = se.ReadRegister(RegRHR, channel)
		if err != nil {
			return
		}

		data = append(data, value)

		xfer++
		avail--
	}

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
