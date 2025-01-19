package gc9a01

func (se *GC9A01) writeCommand(cmd byte) (err error) {

	se.dcPin.SetLow()

	_, err = se.spi.Send([]byte{cmd})

	return
}

func (se *GC9A01) writeCommandBytes(cmd byte, data ...byte) (err error) {

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
	_ = se.writeCommand(MEM_WR)
	err = se.writeBytes(data)
	return
}

func (se *GC9A01) writeMemCont(data []byte) (err error) {
	_ = se.writeCommand(MEM_WR_CONT)
	err = se.writeBytes(data)
	return
}
