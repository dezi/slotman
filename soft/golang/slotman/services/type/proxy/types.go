package proxy

import "slotman/drivers/iface/gpio"

type Area string

const (
	AreaGpio Area = "gpio"
)

type Message struct {
	Area Area
}

type GpioWhat string

const (
	GpioWhatHasGpio   GpioWhat = "gpio.hasGpio"
	GpioWhatOpen      GpioWhat = "gpio.open"
	GpioWhatClose     GpioWhat = "gpio.close"
	GpioWhatSetInput  GpioWhat = "gpio.set.input"
	GpioWhatSetOutput GpioWhat = "gpio.set.output"
	GpioWhatSetLow    GpioWhat = "gpio.set.low"
	GpioWhatSetHigh   GpioWhat = "gpio.set.high"
	GpioWhatGetState  GpioWhat = "gpio.get.State"
)

type Gpio struct {

	//
	// Routing part.
	//

	Area Area
	What GpioWhat

	//
	// Request part.
	//

	PinNo uint8 `json:",omitempty"`

	//
	// Response part.
	//

	State gpio.State `json:",omitempty"`

	Ok  bool  `json:",omitempty"`
	Err error `json:",omitempty"`
}
