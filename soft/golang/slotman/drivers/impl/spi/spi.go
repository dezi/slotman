package spi

func NewDevice(devicePath string) (spi *Device) {
	spi = &Device{Path: devicePath}
	return
}

func (spi *Device) GetDevice() (device string) {
	device = spi.Path
	return
}
