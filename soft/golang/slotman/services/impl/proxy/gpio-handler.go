package proxy

import (
	"encoding/json"
	"slotman/drivers/impl/gpio"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleGpio(reqBytes []byte) (resBytes []byte, err error) {

	sv.gpioDevLock.Lock()
	defer sv.gpioDevLock.Unlock()

	req := proxy.Gpio{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	//
	// Check for calls w/o pin.
	//

	if req.What == proxy.GpioWhatHasGpio {
		req.Ok, req.Err = gpio.HasGpio()
		resBytes, err = json.Marshal(req)
		return
	}

	//
	// Check and create device.
	//

	gpioDev := sv.gpioDevMap[req.PinNo]
	if gpioDev == nil {
		gpioDev = gpio.NewPin(req.PinNo)
		sv.gpioDevMap[req.PinNo] = gpioDev
	}

	switch req.What {

	case proxy.GpioWhatOpen:
		req.Err = gpioDev.Open()
		log.Printf("GPIO Open pin=%d err=%v", gpioDev.GetPinNo(), err)

	case proxy.GpioWhatClose:
		req.Err = gpioDev.Close()
		log.Printf("GPIO Close pin=%d err=%v", gpioDev.GetPinNo(), err)

	case proxy.GpioWhatSetInput:
		req.Err = gpioDev.SetInput()
		log.Printf("GPIO SetInput pin=%d err=%v", gpioDev.GetPinNo(), err)

	case proxy.GpioWhatSetOutput:
		req.Err = gpioDev.SetOutput()
		log.Printf("GPIO SetOutput pin=%d err=%v", gpioDev.GetPinNo(), err)

	case proxy.GpioWhatSetLow:
		req.Err = gpioDev.SetLow()
		//log.Printf("GPIO SetLow  pin=%d err=%v", gpioDev.GetPinNo(), err)

	case proxy.GpioWhatSetHigh:
		req.Err = gpioDev.SetHigh()
		//log.Printf("GPIO SetHigh pin=%d err=%v", gpioDev.GetPinNo(), err)

	case proxy.GpioWhatGetState:
		req.State, req.Err = gpioDev.GetState()
		//log.Printf("GPIO GetState state=%d pin=%d err=%v", req.State, gpioDev.GetPinNo(), err)
	}

	req.Ok = req.Err == nil

	resBytes, err = json.Marshal(req)
	return
}
