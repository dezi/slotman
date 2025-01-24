package proxy

import (
	"slotman/drivers/iface/gpio"
	"slotman/services/impl/provider"
)

const (
	Service provider.Service = "serviceProxy<"
)

type Interface interface {
	GetName() (name provider.Service)

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
	// SPI interface.
	//

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
