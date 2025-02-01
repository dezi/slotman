package proxy

import (
	"encoding/json"
	"errors"
	"slotman/drivers/iface/spi"
	"slotman/services/type/proxy"
)

func (sv *Service) SpiGetDevicePaths() (devicePaths []string, err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatGetDevicePaths,
	}

	res, err := sv.spiExecuteRequest(req, nil)
	if err != nil {
		return
	}

	devicePaths, err = res.Paths, res.NE
	return
}

func (sv *Service) SpiOpen(spi spi.Spi) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatOpen,
	}

	res, err := sv.spiExecuteRequest(req, spi)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) SpiClose(spi spi.Spi) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatClose,
	}

	res, err := sv.spiExecuteRequest(req, spi)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) SpiSetMode(spi spi.Spi, mode uint8) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatSetMode,
		Mode: mode,
	}

	res, err := sv.spiExecuteRequest(req, spi)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) SpiSetBitsPerWord(spi spi.Spi, bpw uint8) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatSetBpw,
		Bpw:  bpw,
	}

	res, err := sv.spiExecuteRequest(req, spi)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) SpiSetSpeed(spi spi.Spi, speed uint32) (err error) {

	req := &proxy.Spi{
		Area:  proxy.AreaSpi,
		What:  proxy.SpiWhatSetSpeed,
		Speed: speed,
	}

	res, err := sv.spiExecuteRequest(req, spi)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) SpiSend(spi spi.Spi, request []byte) (response []byte, err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatSend,
		Send: request,
	}

	res, err := sv.spiExecuteRequest(req, spi)
	if err != nil {
		return
	}

	response, err = res.Recv, res.NE
	return
}

func (sv *Service) spiExecuteRequest(req *proxy.Spi, spi spi.Spi) (res *proxy.Spi, err error) {

	if spi != nil {
		req.Device = spi.GetDevice()
	}

	var resBytes []byte
	resBytes, err = sv.ProxyRequest(req)
	if err != nil {
		return
	}

	res = &proxy.Spi{}
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
