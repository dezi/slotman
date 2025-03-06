package uart

import (
	"slotman/drivers/iface/uart"
	"slotman/utils/log"
	"strings"
)

var (
	//
	// Additional device paths from I2C dual uart devices.
	//
	i2cUarts []uart.Uart
)

func AddI2CUart(i2cUart uart.Uart) {
	i2cUarts = append(i2cUarts, i2cUart)
}

func NewDevice(devicePath string, baudRate int) (uart uart.Uart) {

	if strings.HasPrefix(devicePath, "/dev/i2c") {
		for _, i2cUart := range i2cUarts {
			if i2cUart.GetDevice() == devicePath {
				err := i2cUart.SetBaudrate(baudRate)
				log.Cerror(err)
				return i2cUart
			}
		}
	}

	uart = &Device{
		Path:     devicePath,
		BaudRate: baudRate}
	return
}

func (uart *Device) GetDevice() (device string) {
	device = uart.Path
	return
}

func (uart *Device) SetBaudrate(baudrate int) (err error) {
	uart.BaudRate = baudrate
	return
}

func (uart *Device) GetBaudrate() (baudrate int) {
	baudrate = uart.BaudRate
	return
}
