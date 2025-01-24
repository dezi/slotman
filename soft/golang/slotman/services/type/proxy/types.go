package proxy

type Area string

const (
	AreaGpio Area = "gpio"
)

type Message struct {
	Area Area
}

type GpioWhat string

const (
	GpioWhatHasGpio GpioWhat = "gpio.hasGpio"
	GpioWhatOpen    GpioWhat = "gpio.open"
	GpioWhatClose   GpioWhat = "gpio.close"
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

	Ok  bool  `json:",omitempty"`
	Err error `json:",omitempty"`
}
