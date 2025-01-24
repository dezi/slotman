package spi

import "slotman/services/iface/proxy"

func (spi *Device) Open() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.SpiOpen(spi)
	return
}

func (spi *Device) Close() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.SpiClose(spi)
	return
}

func (spi *Device) SetMode(mode uint8) (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.SpiSetMode(spi, mode)
	return
}

func (spi *Device) SetBitsPerWord(bpw uint8) (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.SpiSetBitsPerWord(spi, bpw)
	return
}

func (spi *Device) SetSpeed(speed uint32) (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.SpiSetSpeed(spi, speed)
	return
}

func (spi *Device) Send(request []byte) (result []byte, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	result, err = prx.SpiSend(spi, request)
	return
}
