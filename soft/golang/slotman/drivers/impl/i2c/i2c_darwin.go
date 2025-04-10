package i2c

import (
	"slotman/services/iface/proxy"
)

// GetDevicePaths Retrieve all I2C device paths.
func GetDevicePaths() (devicePaths []string, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	devicePaths, err = prx.I2cGetDevicePaths()
	return
}

// Open I2C-connection.
func (i2c *Device) Open() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.I2cOpen(i2c)
	return
}

// Close I2C-connection.
func (i2c *Device) Close() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.I2cClose(i2c)
	return
}

func (i2c *Device) TransLock() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.I2cTransLock(i2c)
	return
}

func (i2c *Device) TransUnlock() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	err = prx.I2cTransUnlock(i2c)
	return
}

func (i2c *Device) Write(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	xfer, err = prx.I2cWrite(i2c, data)
	return
}

func (i2c *Device) Read(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	xfer, err = prx.I2cRead(i2c, data)
	return
}

func (i2c *Device) WriteUart(channel byte, timeOut int, data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	xfer, err = prx.I2cWriteUart(i2c, channel, timeOut, data)
	return
}

func (i2c *Device) ReadUart(channel byte, timeOut int, data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	xfer, err = prx.I2cReadUart(i2c, channel, timeOut, data)
	return
}
