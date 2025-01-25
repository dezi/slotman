package uart

import (
	"go.bug.st/serial"
	"time"
)

func GetDevicePaths() (devicePaths []string, err error) {
	devicePaths, err = serial.GetPortsList()
	return
}

func (uart *Device) Open() (err error) {

	uart.port, err = serial.Open(uart.Path, &serial.Mode{
		BaudRate: uart.BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})

	return
}

func (uart *Device) Close() (err error) {
	err = uart.port.Close()
	return
}

func (uart *Device) SetReadTimeout(t time.Duration) (err error) {
	err = uart.port.SetReadTimeout(t)
	return
}

func (uart *Device) Read(p []byte) (n int, err error) {
	n, err = uart.port.Read(p)
	return
}

func (uart *Device) Write(p []byte) (n int, err error) {
	n, err = uart.port.Write(p)
	return
}
