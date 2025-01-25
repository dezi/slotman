package uart

func NewDevice(devicePath string, baudRate int) (uart *Device) {
	uart = &Device{
		Path:     devicePath,
		BaudRate: baudRate}
	return
}

func (uart *Device) GetDevice() (device string) {
	device = uart.Path
	return
}
