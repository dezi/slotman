package sc16is752

func (se *SC15IS752) SetHandler(handler Handler) {
	se.handler = handler
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
