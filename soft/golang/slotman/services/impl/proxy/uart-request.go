package proxy

import (
	"encoding/json"
	"slotman/drivers/iface/uart"
	"slotman/services/type/proxy"
	"time"
)

func (sv *Service) UartGetDevicePaths() (devicePaths []string, err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatGetDevicePaths,
	}

	res, err := sv.uartBuildRequest(req, nil)
	if err != nil {
		return
	}

	devicePaths, err = res.Paths, res.Err
	return
}

func (sv *Service) UartOpen(uart uart.Uart) (err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatOpen,
	}

	res, err := sv.uartBuildRequest(req, uart)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) UartClose(uart uart.Uart) (err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatClose,
	}

	res, err := sv.uartBuildRequest(req, uart)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) UartSetReadTimeout(uart uart.Uart, timeout time.Duration) (err error) {

	req := &proxy.Uart{
		Area:    proxy.AreaUart,
		What:    proxy.UartWhatSetReadTimeout,
		TimeOut: timeout,
	}

	res, err := sv.uartBuildRequest(req, uart)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) UartRead(uart uart.Uart, data []byte) (xfer int, err error) {

	req := &proxy.Uart{
		Area: proxy.AreaUart,
		What: proxy.UartWhatRead,
		Size: len(data),
	}

	res, err := sv.uartBuildRequest(req, uart)
	if err != nil {
		return
	}

	copy(data, req.Read)

	xfer, err = res.Xfer, res.Err
	return
}

func (sv *Service) UartWrite(uart uart.Uart, data []byte) (xfer int, err error) {

	req := &proxy.Uart{
		Area:  proxy.AreaUart,
		What:  proxy.UartWhatWrite,
		Write: data,
	}

	res, err := sv.uartBuildRequest(req, uart)
	if err != nil {
		return
	}

	xfer, err = res.Xfer, res.Err
	return
}

func (sv *Service) uartBuildRequest(req *proxy.Uart, uart uart.Uart) (res *proxy.Uart, err error) {

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
		res = nil
		return
	}

	return
}
