package proxy

import (
	"encoding/json"
	"fmt"
	"slotman/drivers/impl/uart"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleUart(sender string, reqBytes []byte) (resBytes []byte, err error) {

	sv.uartDevLock.Lock()
	defer sv.uartDevLock.Unlock()

	req := proxy.Uart{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	//
	// Check for calls w/o device.
	//

	if req.What == proxy.UartWhatGetDevicePaths {
		req.Paths, req.NE = uart.GetDevicePaths()
		log.Printf("UART GetDevicePaths paths=%v err=%v", req.Paths, req.NE)
		resBytes, err = json.Marshal(req)
		return
	}

	//
	// Check and create device.
	//

	devAddr := fmt.Sprintf("%s-%s", sender, req.Device)

	uartDev := sv.uartDevMap[devAddr]
	if uartDev == nil {
		uartDev = uart.NewDevice(req.Device, req.Baudrate)
		sv.uartDevMap[devAddr] = uartDev
	}

	switch req.What {

	case proxy.UartWhatOpen:
		req.NE = uartDev.Open()
		log.Printf("UART Open dev=%s err=%v", uartDev.GetDevice(), req.NE)

	case proxy.UartWhatClose:
		req.NE = uartDev.Close()
		log.Printf("UART Close dev=%s err=%v", uartDev.GetDevice(), req.NE)

	case proxy.UartWhatSetReadTimeout:
		req.NE = uartDev.SetReadTimeout(req.TimeOut)
		log.Printf("UART SetReadTimeout timeOut=%d dev=%s err=%v", req.TimeOut, uartDev.GetDevice(), req.NE)
		req.TimeOut = 0

	case proxy.UartWhatWrite:
		req.Xfer, req.NE = uartDev.Write(req.Write)
		//log.Printf("UART Write write=%d xfer=%d dev=%s err=%v", len(req.Write), req.Xfer, uartDev.GetDevice(), req.NE)
		req.Write = nil

	case proxy.UartWhatRead:
		req.Read = make([]byte, req.Size)
		req.Xfer, req.NE = uartDev.Read(req.Read)
		req.Read = req.Read[:req.Xfer]
		//log.Printf("UART Read size=%d xfer=%d dev=%s err=%v", req.Size, req.Xfer, uartDev.GetDevice(), req.NE)
		//log.Printf("UART Read [ %02x ]", req.Read)
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
