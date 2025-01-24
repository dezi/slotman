package spi

func NewDevice(devicePath string) (spi *Device) {
	spi = &Device{Path: devicePath}
	return
}
