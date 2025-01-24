package proxy

import (
	"slotman/drivers/iface/gpio"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) GpioHasGpio() (ok bool, err error) {

	log.Printf("############# GpioHasGpio request")

	message := proxy.Gpio{
		Area: proxy.AreaGpio,
		What: proxy.GpioWhatHasGpio,
	}

	err = sv.WriteMessage(message)

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
