package proxy

import (
	"encoding/json"
	"slotman/drivers/impl/uart"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleUart(reqBytes []byte) (resBytes []byte, err error) {

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
		req.Paths, req.Err = uart.GetDevicePaths()
		resBytes, err = json.Marshal(req)
		return
	}

	//
	// Check and create device.
	//

	uartDev := sv.uartDevMap[req.Device]
	if uartDev == nil {
		uartDev = uart.NewDevice(req.Device, req.Baudrate)
		sv.uartDevMap[req.Device] = uartDev
	}

	switch req.What {

	case proxy.UartWhatOpen:
		req.Err = uartDev.Open()
		log.Printf("UART Open dev=%s err=%v", uartDev.GetDevice(), err)

	case proxy.UartWhatClose:
		req.Err = uartDev.Close()
		log.Printf("UART Close dev=%s err=%v", uartDev.GetDevice(), err)

	case proxy.UartWhatSetReadTimeout:
		req.Err = uartDev.SetReadTimeout(req.TimeOut)
		log.Printf("UART SetReadTimeout timeOut=%d dev=%s err=%v", req.TimeOut, uartDev.GetDevice(), err)
		req.TimeOut = 0

	case proxy.UartWhatWrite:
		req.Xfer, req.Err = uartDev.Write(req.Write)
		log.Printf("UART Write write=%d dev=%s err=%v", len(req.Write), uartDev.GetDevice(), err)
		req.Write = nil

	case proxy.UartWhatRead:
		req.Read = make([]byte, req.Size)
		req.Xfer, req.Err = uartDev.Read(req.Read)
		log.Printf("UART Read read=%d dev=%s err=%v", len(req.Read), uartDev.GetDevice(), err)
	}

	req.Ok = req.Err == nil

	resBytes, err = json.Marshal(req)

	return
}
