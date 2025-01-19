package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"slotman/utils/log"
	"time"
)

type State uint8

const (
	Low State = iota
	High
)

type Pin struct {
	pin     rpio.Pin
	on      time.Duration
	off     time.Duration
	loops   int
	started bool
}

var (
	isOpen bool
)

func HasGpio() (ok bool) {
	_, tryErr := os.Stat("/dev/gpiomem")
	ok = tryErr == nil
	return
}

func GetPin(pinNo uint8) (pin *Pin, err error) {

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

	pin = &Pin{
		pin: rpio.Pin(pinNo),
	}

	return
}

func (pin *Pin) SetOutput() {
	pin.pin.Output()
}

func (pin *Pin) SetInput() {
	pin.pin.Input()
}

func (pin *Pin) SetLow() {
	pin.pin.Low()
}

func (pin *Pin) SetHigh() {
	pin.pin.High()
}

func (pin *Pin) Read() (state State) {
	state = State(pin.pin.Read())
	return
}
