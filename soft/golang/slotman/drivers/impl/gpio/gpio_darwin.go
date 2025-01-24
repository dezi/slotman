package gpio

import gpio2 "slotman/drivers/types/gpio"

func HasGpio() (ok bool) {

	//prx, err := proxy.GetInstance()
	//if err != nil {
	//	return
	//}

	//_ = prx
	//ok, err = prx.GpioHasGpio()
	return
}

func (pin *gpio2.Pin) Open() (err error) {
	return
}

func (pin *gpio2.Pin) Close() (err error) {
	return
}

func (pin *gpio2.Pin) SetOutput() (err error) {
	return
}

func (pin *gpio2.Pin) SetInput() (err error) {
	return
}

func (pin *gpio2.Pin) SetLow() (err error) {
	return
}

func (pin *gpio2.Pin) SetHigh() (err error) {
	return
}

func (pin *gpio2.Pin) GetState() (state gpio2.State, err error) {
	return
}
