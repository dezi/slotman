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

	log.Printf("########### reqBytes=%s", string(reqBytes))
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

	case proxy.GpioWhatClose:
		req.Err = gpioDev.Close()

	case proxy.GpioWhatSetInput:
		req.Err = gpioDev.SetInput()

	case proxy.GpioWhatSetOutput:
		req.Err = gpioDev.SetOutput()

	case proxy.GpioWhatSetLow:
		log.Printf("Gpio set low  pin=%d %d", gpioDev.GetPinNo(), req.PinNo)
		req.Err = gpioDev.SetLow()

	case proxy.GpioWhatSetHigh:
		log.Printf("Gpio set high pin=%d %d", gpioDev.GetPinNo(), req.PinNo)
		req.Err = gpioDev.SetHigh()

	case proxy.GpioWhatGetState:
		req.State, req.Err = gpioDev.GetState()
	}

	req.Ok = req.Err == nil

	resBytes, err = json.Marshal(req)
	return
}
