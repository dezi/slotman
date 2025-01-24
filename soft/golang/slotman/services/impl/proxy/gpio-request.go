package proxy

import (
	"encoding/json"
	"slotman/drivers/iface/gpio"
	"slotman/services/type/proxy"
)

func (sv *Service) GpioHasGpio() (ok bool, err error) {

	req := proxy.Gpio{
		Area: proxy.AreaGpio,
		What: proxy.GpioWhatHasGpio,
	}

	var resBytes []byte
	resBytes, err = sv.ProxyRequest(req)
	if err != nil {
		return
	}

	res := proxy.Gpio{}
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		return
	}

	ok, err = res.Ok, res.Err
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
