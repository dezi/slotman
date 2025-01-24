package uart

import (
	"go.bug.st/serial"
	"time"
)

type Uart struct {
	Device   string
	BaudRate int

	port serial.Port
}

func GetPortsList() (ports []string, err error) {

	ports, err = serial.GetPortsList()
	if err != nil {
		return
	}

	return
}

func Open(device string, baudRate int) (uart *Uart, err error) {

	uart = &Uart{
		Device:   device,
		BaudRate: baudRate,
	}

	uart.port, err = serial.Open(device, &serial.Mode{
		BaudRate: baudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})

	if err != nil {
		uart = nil
	}

	return
}

func (uart Uart) Close() (err error) {
	err = uart.port.Close()
	return
}

func (uart Uart) SetReadTimeout(t time.Duration) (err error) {
	err = uart.port.SetReadTimeout(t)
	return
}

func (uart Uart) Read(p []byte) (n int, err error) {
	n, err = uart.port.Read(p)
	return
}

func (uart Uart) Write(p []byte) (n int, err error) {
	n, err = uart.port.Write(p)
	return
}
