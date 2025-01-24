package proxy

import (
	"encoding/json"
	"slotman/drivers/iface/gpio"
	"slotman/services/type/proxy"
)

func (sv *Service) GpioHasGpio() (ok bool, err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatHasGpio, nil)
	if err != nil {
		return
	}

	ok, err = res.Ok, res.Err
	return
}

func (sv *Service) GpioOpen(pin gpio.Gpio) (err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatOpen, nil)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) GpioClose(pin gpio.Gpio) (err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatClose, nil)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) GpioSetOutput(pin gpio.Gpio) (err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatSetOutput, nil)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) GpioSetInput(pin gpio.Gpio) (err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatSetInput, nil)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) GpioSetLow(pin gpio.Gpio) (err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatSetLow, nil)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) GpioSetHigh(pin gpio.Gpio) (err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatSetHigh, nil)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) GpioGetState(pin gpio.Gpio) (state gpio.State, err error) {

	res, err := sv.gpioBuildRequest(proxy.GpioWhatGetState, nil)
	if err != nil {
		return
	}

	state, err = res.State, res.Err
	return
}

func (sv *Service) gpioBuildRequest(what proxy.GpioWhat, pin gpio.Gpio) (res *proxy.Gpio, err error) {

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

	return
}
