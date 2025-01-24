package proxy

type Area string

const (
	AreaGpio Area = "gpio"
)

type Message struct {
	Area Area
	What string
	Path string
}

type GpioWhat string

const (
	GpioWhatHasGpio GpioWhat = "gpio.hasGpio"
)

type Gpio struct {
	Area Area
	What GpioWhat

	PinNo string `json:",omitempty"`

	Ok bool `json:",omitempty"`
}
