package proxy

import (
	gpio2 "slotman/drivers/impl/gpio"
)

func (sv *Service) GpioHasGpio() (ok bool, err error) {

	target, err := sv.getTarget()
	if err != nil {
		return
	}

	_ = target

	return
}

func (sv *Service) GpioOpen(pin gpio2.Gpio) (err error) {
	return
}

func (sv *Service) GpioClose(pin gpio2.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetOutput(pin gpio2.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetInput(pin gpio2.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetLow(pin gpio2.Gpio) (err error) {
	return
}

func (sv *Service) GpioSetHigh(pin gpio2.Gpio) (err error) {
	return
}

func (sv *Service) GpioGetState(pin gpio2.Gpio) (state gpio2.State, err error) {
	return
}
