package proxy

import (
	"encoding/json"
	"errors"
	"slotman/drivers/iface/i2c"
	"slotman/services/type/proxy"
)

func (sv *Service) I2cGetDevicePaths() (devicePaths []string, err error) {

	req := &proxy.I2c{
		Area: proxy.AreaI2c,
		What: proxy.I2cWhatGetDevicePaths,
	}

	res, err := sv.i2cExecuteRequest(req, nil)
	if err != nil {
		return
	}

	devicePaths, err = res.Paths, res.NE
	return
}

func (sv *Service) I2cOpen(i2c i2c.I2c) (err error) {

	req := &proxy.I2c{
		Area: proxy.AreaI2c,
		What: proxy.I2cWhatOpen,
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) I2cClose(i2c i2c.I2c) (err error) {

	req := &proxy.I2c{
		Area: proxy.AreaI2c,
		What: proxy.I2cWhatClose,
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) I2cTransLock(i2c i2c.I2c) (err error) {

	req := &proxy.I2c{
		Area: proxy.AreaI2c,
		What: proxy.I2cWhatTransLock,
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) I2cTransUnlock(i2c i2c.I2c) (err error) {

	req := &proxy.I2c{
		Area: proxy.AreaI2c,
		What: proxy.I2cWhatTransUnlock,
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	err = res.NE
	return
}

func (sv *Service) I2cWrite(i2c i2c.I2c, data []byte) (xfer int, err error) {

	req := &proxy.I2c{
		Area:  proxy.AreaI2c,
		What:  proxy.I2cWhatWrite,
		Write: data,
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	xfer, err = res.Xfer, res.NE
	return
}

func (sv *Service) I2cRead(i2c i2c.I2c, data []byte) (xfer int, err error) {

	req := &proxy.I2c{
		Area: proxy.AreaI2c,
		What: proxy.I2cWhatRead,
		Size: len(data),
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	copy(data, res.Read)

	xfer, err = res.Xfer, res.NE
	return
}

func (sv *Service) I2cReadUart(i2c i2c.I2c, channel byte, timeOut int, data []byte) (xfer int, err error) {

	req := &proxy.I2c{
		Area:    proxy.AreaI2c,
		What:    proxy.I2cWhatReadUart,
		Size:    len(data),
		Channel: channel,
		TimeOut: timeOut,
	}

	res, err := sv.i2cExecuteRequest(req, i2c)
	if err != nil {
		return
	}

	copy(data, res.Read)

	xfer, err = res.Xfer, res.NE
	return
}

func (sv *Service) i2cExecuteRequest(req *proxy.I2c, i2c i2c.I2c) (res *proxy.I2c, err error) {

	if i2c != nil {
		req.Device = i2c.GetDevice()
		req.Addr = i2c.GetAddr()
	}

	var resBytes []byte
	resBytes, err = sv.ProxyRequest(req)
	if err != nil {
		return
	}

	res = &proxy.I2c{}
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
