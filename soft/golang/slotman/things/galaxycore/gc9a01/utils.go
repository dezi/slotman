package gc9a01

func (se *GC9A01) writeCommand(cmd Command) (err error) {

	se.dcPin.SetLow()

	_, err = se.spi.Send([]byte{byte(cmd)})

	return
}

func (se *GC9A01) writeCommandBytes(cmd Command, data ...byte) (err error) {

	_ = se.writeCommand(cmd)
	err = se.writeBytes(data)

	return
}

func (se *GC9A01) writeByte(data byte) (err error) {

	se.dcPin.SetHigh()

	_, err = se.spi.Send([]byte{data})

	return
}

func (se *GC9A01) writeBytes(data []byte) (err error) {

	se.dcPin.SetHigh()

	_, err = se.spi.Send(data)

	return
}

func (se *GC9A01) writeMem(data []byte) (err error) {
	_ = se.writeCommand(CommandMemWr)
	err = se.writeBytes(data)
	return
}

func (se *GC9A01) writeMemCont(data []byte) (err error) {
	_ = se.writeCommand(CommandMemWrCont)
	err = se.writeBytes(data)
	return
}
