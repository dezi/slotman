package proxy

import (
	"encoding/json"
	"errors"
	"slotman/drivers/iface/uart"
	"slotman/services/type/proxy"
	"slotman/utils/log"
	"time"
)

func (sv *Service) UartGetDevicePaths() (devicePaths []string, err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatGetDevicePaths,
	}

	res, err := sv.uartExecuteRequest(req, nil)
	if err != nil {
		return
	}

	devicePaths, err = res.Paths, res.NE
	return
}

func (sv *Service) UartOpen(uart uart.Uart) (err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatOpen,
	}

	res, err := sv.uartExecuteRequest(req, uart)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) UartClose(uart uart.Uart) (err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatClose,
	}

	res, err := sv.uartExecuteRequest(req, uart)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) UartSetReadTimeout(uart uart.Uart, timeout time.Duration) (err error) {

	req := &proxy.Uart{
		Area:    proxy.AreaUart,
		What:    proxy.UartWhatSetReadTimeout,
		TimeOut: timeout,
	}

	res, err := sv.uartExecuteRequest(req, uart)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) UartRead(uart uart.Uart, data []byte) (xfer int, err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatRead,
		Size: len(data),
	}

	res, err := sv.uartExecuteRequest(req, uart)
	if err != nil {
		return
	}

	copy(data, res.Read)

	xfer, err = res.Xfer, res.NE
	return
}

func (sv *Service) UartWrite(uart uart.Uart, data []byte) (xfer int, err error) {

	req := &proxy.Uart{
		Area:  proxy.AreaUart,
		What:  proxy.UartWhatWrite,
		Write: data,
	}

	res, err := sv.uartExecuteRequest(req, uart)
	if err != nil {
		return
	}

	xfer, err = res.Xfer, res.NE
	return
}

func (sv *Service) uartExecuteRequest(req *proxy.Uart, uart uart.Uart) (res *proxy.Uart, err error) {

	if uart != nil {
		req.Device = uart.GetDevice()
		req.Baudrate = uart.GetBaudrate()
	}

	var resBytes []byte
	resBytes, err = sv.proxyRequest(req)
	if err != nil {
		return
	}

	res = &proxy.Uart{}
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		log.Printf("########## resBytes=%s", string(resBytes))
		res = nil
		return
	}

	if res.Err != "" {
		res.NE = errors.New(res.Err)
	}

	return
}
