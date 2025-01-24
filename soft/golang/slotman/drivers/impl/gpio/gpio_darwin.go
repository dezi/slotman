package gpio

import (
	"slotman/drivers/iface/gpio"
	"slotman/services/iface/proxy"
)

func HasGpio() (ok bool, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	ok, err = prx.GpioHasGpio()
	return
}

func (pin *Pin) Open() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.GpioOpen(pin)
	return
}

func (pin *Pin) Close() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.GpioClose(pin)
	return
}

func (pin *Pin) SetOutput() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.GpioSetOutput(pin)
	return
}

func (pin *Pin) SetInput() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.GpioSetInput(pin)
	return
}

func (pin *Pin) SetLow() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.GpioSetLow(pin)
	return
}

func (pin *Pin) SetHigh() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.GpioSetHigh(pin)
	return
}

func (pin *Pin) GetState() (state gpio.State, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	state, err = prx.GpioGetState(pin)
	return
}
