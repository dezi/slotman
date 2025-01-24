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

	PinNo string `json:",omitempty"`

	//
	// Response part.
	//

	Ok bool `json:",omitempty"`
}
