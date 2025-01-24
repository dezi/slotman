package proxy

import (
	"encoding/json"
	"slotman/drivers/iface/spi"
	"slotman/services/type/proxy"
)

func (sv *Service) SpiGetDevicePaths() (devicePaths []string, err error) {
	return
}

func (sv *Service) SpiOpen(spi spi.Spi) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatOpen,
	}

	res, err := sv.spiBuildRequest(req, spi)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) SpiClose(spi spi.Spi) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatClose,
	}

	res, err := sv.spiBuildRequest(req, spi)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) SpiSetMode(spi spi.Spi, mode uint8) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatSetMode,
		Mode: mode,
	}

	res, err := sv.spiBuildRequest(req, spi)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) SpiSetBitsPerWord(spi spi.Spi, bpw uint8) (err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatSetBpw,
		Bpw:  bpw,
	}

	res, err := sv.spiBuildRequest(req, spi)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) SpiSetSpeed(spi spi.Spi, speed uint32) (err error) {

	req := &proxy.Spi{
		Area:  proxy.AreaSpi,
		What:  proxy.SpiWhatSetSpeed,
		Speed: speed,
	}

	res, err := sv.spiBuildRequest(req, spi)
	if err != nil {
		return
	}

	err = res.Err
	return
}

func (sv *Service) SpiSend(spi spi.Spi, request []byte) (result []byte, err error) {

	req := &proxy.Spi{
		Area: proxy.AreaSpi,
		What: proxy.SpiWhatSend,
		Send: request,
	}

	res, err := sv.spiBuildRequest(req, spi)
	if err != nil {
		return
	}

	result, err = res.Recv, res.Err
	return
}

func (sv *Service) spiBuildRequest(req *proxy.Spi, spi spi.Spi) (res *proxy.Spi, err error) {

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

	return
}
