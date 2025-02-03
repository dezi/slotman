package sc16is752

func (se *SC15IS752) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *SC15IS752) SetPollInterval(channel byte, millis int) (err error) {

	if channel > ChannelB {
		err = ErrInvalidChannel
		return
	}

	se.pollSleep[channel] = millis
	return
}

func (se *SC15IS752) EnableFifo(channel byte, enable bool) (err error) {

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
