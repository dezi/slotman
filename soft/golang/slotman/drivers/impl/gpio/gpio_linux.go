package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"slotman/drivers/iface/gpio"
	"slotman/utils/log"
)

var (
	isOpen bool
)

func HasGpio() (ok bool, err error) {
	_, tryErr := os.Stat("/dev/gpiomem")
	ok = tryErr == nil
	return
}

func (pin *Pin) Open() (err error) {

	if !isOpen {

		err = rpio.Open()
		if err != nil {
			log.Cerror(err)
			return
		}

		isOpen = true
	}

	_, err = os.Stat("/dev/gpiomem")
	if err != nil {
		log.Cerror(err)
		return
	}

	pin.pin = rpio.Pin(pin.PinNo)
	return
}

func (pin *Pin) Close() (err error) {
	return
}

func (pin *Pin) SetOutput() (err error) {
	pin.pin.Output()
	return
}

func (pin *Pin) SetInput() (err error) {
	pin.pin.Input()
	return
}

func (pin *Pin) SetLow() (err error) {
	pin.pin.Low()
	return
}

func (pin *Pin) SetHigh() (err error) {
	pin.pin.High()
	return
}

func (pin *Pin) GetState() (state gpio.State, err error) {
	state = gpio.State(pin.pin.Read())
	return
}
