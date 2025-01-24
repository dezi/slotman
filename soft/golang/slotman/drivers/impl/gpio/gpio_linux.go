package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	gpio2 "slotman/drivers/types/gpio"
	"slotman/utils/log"
)

var (
	isOpen bool
)

func HasGpio() (ok bool) {
	_, tryErr := os.Stat("/dev/gpiomem")
	ok = tryErr == nil
	return
}

func (pin *gpio2.Pin) Open() (err error) {

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

	pin.pin = rpio.Pin(pin.pinNo)
	return
}

func (pin *gpio2.Pin) Close() (err error) {
	return
}

func (pin *gpio2.Pin) SetOutput() (err error) {
	pin.pin.Output()
	return
}

func (pin *gpio2.Pin) SetInput() (err error) {
	pin.pin.Input()
	return
}

func (pin *gpio2.Pin) SetLow() (err error) {
	pin.pin.Low()
	return
}

func (pin *gpio2.Pin) SetHigh() (err error) {
	pin.pin.High()
	return
}

func (pin *gpio2.Pin) GetState() (state gpio2.State, err error) {
	state = gpio2.State(pin.pin.Read())
	return
}
