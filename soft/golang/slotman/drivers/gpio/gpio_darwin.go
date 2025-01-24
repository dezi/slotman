package gpio

import "slotman/services/iface/proxy"

func HasGpio() (ok bool) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	ok, err = prx.GpioHasGpio()
	return
}

func (pin *Pin) Open() (err error) {
	return
}

func (pin *Pin) Close() (err error) {
	return
}

func (pin *Pin) SetOutput() (err error) {
	return
}

func (pin *Pin) SetInput() (err error) {
	return
}

func (pin *Pin) SetLow() (err error) {
	return
}

func (pin *Pin) SetHigh() (err error) {
	return
}

func (pin *Pin) GetState() (state State, err error) {
	return
}
