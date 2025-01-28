package proxy

import (
	"slotman/drivers/iface/gpio"
	"slotman/utils/simple"
)

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

	Uuid simple.UUIDHex

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

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}

func (px *Gpio) GetUuid() (uuid simple.UUIDHex) {
	uuid = px.Uuid
	return
}

func (px *Gpio) SetUuid(uuid simple.UUIDHex) {
	px.Uuid = uuid
	return
}

func (px *Gpio) GetArea() (area Area) {
	area = px.Area
	return
}
