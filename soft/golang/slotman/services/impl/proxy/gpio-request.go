package proxy

import (
	"encoding/json"
	"errors"
	"slotman/drivers/iface/gpio"
	"slotman/services/type/proxy"
)

func (sv *Service) GpioHasGpio() (ok bool, err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatHasGpio, nil)
	if err != nil {
		return
	}

	ok, err = res.Ok, res.NE
	return
}

func (sv *Service) GpioOpen(pin gpio.Gpio) (err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatOpen, pin)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) GpioClose(pin gpio.Gpio) (err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatClose, pin)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) GpioSetOutput(pin gpio.Gpio) (err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatSetOutput, pin)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) GpioSetInput(pin gpio.Gpio) (err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatSetInput, pin)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) GpioSetLow(pin gpio.Gpio) (err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatSetLow, pin)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) GpioSetHigh(pin gpio.Gpio) (err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatSetHigh, pin)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) GpioGetState(pin gpio.Gpio) (state gpio.State, err error) {

	res, err := sv.gpioExecuteRequest(proxy.GpioWhatGetState, pin)
	if err != nil {
		return
	}

	state, err = res.State, res.NE
	return
}

func (sv *Service) gpioExecuteRequest(what proxy.GpioWhat, pin gpio.Gpio) (res *proxy.Gpio, err error) {

	req := &proxy.Gpio{
		Area: proxy.AreaGpio,
		What: what,
	}

	if pin != nil {
		req.PinNo = pin.GetPinNo()
	}

	var resBytes []byte
	resBytes, err = sv.ProxyRequest(req)
	if err != nil {
		return
	}

	res = &proxy.Gpio{}
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		res = nil
		return
	}

	if res.Err != "" {
		res.NE = errors.New(res.Err)
	}

	return
}
