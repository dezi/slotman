package spi

func (spi *Device) Open() (err error) {
	return
}

func (spi *Device) Close() (err error) {
	return
}

func (spi *Device) SetMode(mode uint8) (err error) {
	return
}

func (spi *Device) SetBitsPerWord(bpw uint8) (err error) {
	return
}

func (spi *Device) SetSpeed(speed uint32) (err error) {
	return
}

func (spi *Device) Send(request []byte) (result []byte, err error) {
	return
}
