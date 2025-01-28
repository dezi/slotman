package proxy

import (
	"encoding/json"
	"fmt"
	"slotman/drivers/impl/gpio"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleGpio(sender string, reqBytes []byte) (resBytes []byte, err error) {

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
		req.Ok, req.NE = gpio.HasGpio()
		log.Printf("GPIO HasGpio ok=%v err=%v", req.Ok, req.NE)

		if req.NE == nil {
			req.Ok = true
		} else {
			req.Ok = false
			req.Err = req.NE.Error()
		}

		resBytes, err = json.Marshal(req)
		return
	}

	//
	// Check and create device.
	//

	devAddr := fmt.Sprintf("%s-%d", sender, req.PinNo)

	gpioDev := sv.gpioDevMap[devAddr]
	if gpioDev == nil {
		gpioDev = gpio.NewPin(req.PinNo)
		sv.gpioDevMap[devAddr] = gpioDev
	}

	switch req.What {

	case proxy.GpioWhatOpen:
		req.NE = gpioDev.Open()
		log.Printf("GPIO Open pin=%d err=%v", gpioDev.GetPinNo(), req.NE)

	case proxy.GpioWhatClose:
		req.NE = gpioDev.Close()
		log.Printf("GPIO Close pin=%d err=%v", gpioDev.GetPinNo(), req.NE)

	case proxy.GpioWhatSetInput:
		req.NE = gpioDev.SetInput()
		log.Printf("GPIO SetInput pin=%d err=%v", gpioDev.GetPinNo(), req.NE)

	case proxy.GpioWhatSetOutput:
		req.NE = gpioDev.SetOutput()
		log.Printf("GPIO SetOutput pin=%d err=%v", gpioDev.GetPinNo(), req.NE)

	case proxy.GpioWhatSetLow:
		req.NE = gpioDev.SetLow()
		//log.Printf("GPIO SetLow  pin=%d err=%v", gpioDev.GetPinNo(), req.NE)

	case proxy.GpioWhatSetHigh:
		req.NE = gpioDev.SetHigh()
		//log.Printf("GPIO SetHigh pin=%d err=%v", gpioDev.GetPinNo(), req.NE)

	case proxy.GpioWhatGetState:
		req.State, req.NE = gpioDev.GetState()
		//log.Printf("GPIO GetState state=%d pin=%d err=%v", req.State, gpioDev.GetPinNo(), req.NE)
	}

	if req.NE == nil {
		req.Ok = true
	} else {
		req.Ok = false
		req.Err = req.NE.Error()
	}

	resBytes, err = json.Marshal(req)
	return
}
