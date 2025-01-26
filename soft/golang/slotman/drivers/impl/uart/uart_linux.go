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
	err = uart.port.Close()
	return
}

func (uart *Device) SetReadTimeout(timeout time.Duration) (err error) {
	err = uart.port.SetReadTimeout(timeout)
	return
}

func (uart *Device) Read(data []byte) (xfer int, err error) {
	xfer, err = uart.port.Read(data)
	return
}

func (uart *Device) Write(data []byte) (xfer int, err error) {
	xfer, err = uart.port.Write(data)
	return
}
