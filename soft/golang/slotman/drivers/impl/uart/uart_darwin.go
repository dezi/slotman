package uart

import (
	"go.bug.st/serial"
	"slotman/services/iface/proxy"
	"time"
)

func GetDevicePaths() (devicePaths []string, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	if prx.CheckTarget() {
		devicePaths, err = prx.UartGetDevicePaths()
		return
	}

	devicePaths, err = serial.GetPortsList()
	return
}

func (uart *Device) Open() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	if prx.CheckTarget() {
		err = prx.UartOpen(uart)
		return
	}

	port, err := serial.Open(uart.Path, &serial.Mode{
		BaudRate: uart.BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})

	if err == nil {
		uart.port = port
	}

	return
}

func (uart *Device) Close() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	if prx.CheckTarget() {
		err = prx.UartClose(uart)
		return
	}

	err = uart.port.Close()
	return
}

func (uart *Device) SetReadTimeout(timeout time.Duration) (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	if prx.CheckTarget() {
		err = prx.UartSetReadTimeout(uart, timeout)
		return
	}

	err = uart.port.SetReadTimeout(timeout)
	return
}

func (uart *Device) Read(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	if prx.CheckTarget() {
		xfer, err = prx.UartRead(uart, data)
		return
	}

	xfer, err = uart.port.Read(data)
	return
}

func (uart *Device) Write(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	if prx.CheckTarget() {
		xfer, err = prx.UartWrite(uart, data)
		return
	}

	xfer, err = uart.port.Write(data)
	return
}
