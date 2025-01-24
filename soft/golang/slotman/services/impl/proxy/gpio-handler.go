package proxy

import (
	"encoding/json"
	"slotman/drivers/impl/gpio"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleGpio(reqBytes []byte) (resBytes []byte, err error) {

	req := proxy.Gpio{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	switch req.What {
	case proxy.GpioWhatHasGpio:
		req.Ok, req.Err = gpio.HasGpio()
	}

	resBytes, err = json.Marshal(req)
	return
}
