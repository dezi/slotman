package proxy

import (
	gpio2 "slotman/drivers/impl/gpio"
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

	GpioOpen(pin gpio2.Gpio) (err error)
	GpioClose(pin gpio2.Gpio) (err error)

	GpioSetOutput(pin gpio2.Gpio) (err error)
	GpioSetInput(pin gpio2.Gpio) (err error)
	GpioSetLow(pin gpio2.Gpio) (err error)
	GpioSetHigh(pin gpio2.Gpio) (err error)

	GpioGetState(pin gpio2.Gpio) (state gpio2.State, err error)
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
