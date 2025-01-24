package proxy

import (
	"slotman/drivers/iface/gpio"
)

func (sv *Service) GpioHasGpio() (ok bool, err error) {

	target, err := sv.getTarget()
	if err != nil {
		return
	}

	_ = target

	return
}

func (sv *Service) GpioOpen(pin gpio.Gpio) (err error) {
	return
}

func (sv *Service) GpioClose(pin gpio.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetOutput(pin gpio.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetInput(pin gpio.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetLow(pin gpio.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetHigh(pin gpio.Gpio) (err error) {
	return
}

func (sv *Service) GpioGetState(pin gpio.Gpio) (state gpio.State, err error) {
	return
}
