package sc16is752

import "slotman/things"

func (se *SC15IS752) ReadRegister(register, channel byte) (value byte, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	value, err = se.i2cDev.ReadRegByte(register<<3 | channel<<1)
	return
}

func (se *SC15IS752) WriteRegister(register, channel, value byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	err = se.i2cDev.WriteRegByte(register<<3|channel<<1, value)
	return
}
