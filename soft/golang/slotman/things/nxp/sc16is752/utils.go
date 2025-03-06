package sc16is752

import "slotman/things"

func (se *SC15IS752) ReadRegister(register, channel byte) (value byte, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	value, err = se.i2cDev.ReadRegByte(register<<3 | channel<<1)
	return
}

func (se *SC15IS752) ReadRegisterBytes(register, channel byte, size int) (data []byte, xfer int, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	data, xfer, err = se.i2cDev.ReadRegBytes(register<<3|channel<<1, size)
	return
}

func (se *SC15IS752) WriteRegister(register, channel, value byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	err = se.i2cDev.WriteRegByte(register<<3|channel<<1, value)
	return
}

func (se *SC15IS752) WriteRegisterBytes(register, channel byte, data []byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	err = se.i2cDev.WriteRegBytes(register<<3|channel<<1, data)
	return
}
