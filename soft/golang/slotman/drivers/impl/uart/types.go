package uart

import (
	"go.bug.st/serial"
	"slotman/drivers/iface/uart"
)

type Device struct {
	Path     string
	BaudRate int

	//
	// Local serial port.
	//

	port serial.Port

	//
	// Local or remote I2C serial ports.
	//

	i2cPort uart.Uart
}
