package proxy

import "slotman/drivers/iface/gpio"

type Area string

const (
	AreaGpio Area = "gpio"
	AreaSpi  Area = "spi"
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

type SpiWhat string

const (
	SpiWhatOpen     SpiWhat = "spi.open"
	SpiWhatClose    SpiWhat = "spi.close"
	SpiWhatSetMode  SpiWhat = "spi.set.mode"
	SpiWhatSetBpw   SpiWhat = "spi.set.bpw"
	SpiWhatSetSpeed SpiWhat = "spi.set.speed"
	SpiWhatSend     SpiWhat = "spi.send"
)

type Spi struct {

	//
	// Routing part.
	//

	Area Area
	What SpiWhat

	//
	// Request part.
	//

	Device string `json:",omitempty"`

	//
	// Response part.
	//

	Mode  uint8  `json:",omitempty"`
	Bpw   uint8  `json:",omitempty"`
	Speed uint32 `json:",omitempty"`

	Paths []string `json:",omitempty"`
	Send  []byte   `json:",omitempty"`
	Recv  []byte   `json:",omitempty"`

	Ok  bool  `json:",omitempty"`
	Err error `json:",omitempty"`
}
