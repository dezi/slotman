package proxy

import (
	"slotman/drivers/iface/gpio"
	"slotman/drivers/iface/i2c"
	"slotman/drivers/iface/spi"
	"slotman/drivers/iface/uart"
	"slotman/services/impl/provider"
	"slotman/services/type/proxy"
	"time"
)

const (
	Service provider.Service = "serviceProxy<"
)

type Interface interface {
	GetName() (name provider.Service)

	CheckTarget() (ok bool)

	Subscribe(area proxy.Area, handler proxy.Subscriber)
	Unsubscribe(area proxy.Area)

	ProxyRequest(req proxy.Message) (res []byte, err error)
	ProxyBroadcast(resBytes []byte) (err error)

	//
	// GPIO interface.
	//

	GpioHasGpio() (ok bool, err error)

	GpioOpen(pin gpio.Gpio) (err error)
	GpioClose(pin gpio.Gpio) (err error)

	GpioSetOutput(pin gpio.Gpio) (err error)
	GpioSetInput(pin gpio.Gpio) (err error)
	GpioSetLow(pin gpio.Gpio) (err error)
	GpioSetHigh(pin gpio.Gpio) (err error)

	GpioGetState(pin gpio.Gpio) (state gpio.State, err error)

	//
	// I2C interface.
	//

	I2cGetDevicePaths() (devicePaths []string, err error)

	I2cOpen(i2c i2c.I2c) (err error)
	I2cClose(i2c i2c.I2c) (err error)

	I2cTransLock(i2c i2c.I2c) (err error)
	I2cTransUnlock(i2c i2c.I2c) (err error)

	I2cWrite(i2c i2c.I2c, data []byte) (xfer int, err error)
	I2cRead(i2c i2c.I2c, data []byte) (xfer int, err error)

	//
	// SPI interface.
	//

	SpiGetDevicePaths() (devicePaths []string, err error)

	SpiOpen(spi spi.Spi) (err error)
	SpiClose(spi spi.Spi) (err error)

	SpiSetMode(spi spi.Spi, mode uint8) (err error)
	SpiSetBitsPerWord(spi spi.Spi, bpw uint8) (err error)
	SpiSetSpeed(spi spi.Spi, speed uint32) (err error)

	SpiSend(spi spi.Spi, request []byte) (response []byte, err error)

	//
	// UART interface.
	//

	UartGetDevicePaths() (devicePaths []string, err error)

	UartOpen(uart uart.Uart) (err error)
	UartClose(uart uart.Uart) (err error)

	UartSetReadTimeout(uart uart.Uart, timeout time.Duration) (err error)

	UartRead(uart uart.Uart, data []byte) (xfer int, err error)
	UartWrite(uart uart.Uart, data []byte) (xfer int, err error)
}

func GetInstance() (iface Interface, err error) {

	baseProvider, err := provider.GetProvider(Service)
	if err != nil {
		return
	}

	iface = baseProvider.(Interface)
	if iface == nil {
		err = provider.ErrNotFound(Service)
		return
	}

	return
}
