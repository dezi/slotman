package uart

import (
	"slotman/services/iface/proxy"
	"time"
)

func GetDevicePaths() (devicePaths []string, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	devicePaths, err = prx.UartGetDevicePaths()
	return
}

func (uart *Device) Open() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.UartOpen(uart)
	return
}

func (uart *Device) Close() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.UartClose(uart)
	return
}

func (uart *Device) SetReadTimeout(timeout time.Duration) (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.UartSetReadTimeout(uart, timeout)
	return
}

func (uart *Device) Read(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	xfer, err = prx.UartRead(uart, data)
	return
}

func (uart *Device) Write(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	xfer, err = prx.UartWrite(uart, data)
	return
}
