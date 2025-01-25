package proxy

import (
	"slotman/drivers/iface/gpio"
	"time"
)

type Area string

const (
	AreaGpio Area = "gpio"
	AreaSpi  Area = "spi"
	AreaUart Area = "uart"
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
	SpiWhatGetDevicePaths SpiWhat = "spi.get.device-paths"
	SpiWhatOpen           SpiWhat = "spi.open"
	SpiWhatClose          SpiWhat = "spi.close"
	SpiWhatSetMode        SpiWhat = "spi.set.mode"
	SpiWhatSetBpw         SpiWhat = "spi.set.bpw"
	SpiWhatSetSpeed       SpiWhat = "spi.set.speed"
	SpiWhatSend           SpiWhat = "spi.send"
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

	Mode  uint8  `json:",omitempty"`
	Bpw   uint8  `json:",omitempty"`
	Speed uint32 `json:",omitempty"`

	//
	// Response part.
	//

	Paths []string `json:",omitempty"`
	Send  []byte   `json:",omitempty"`
	Recv  []byte   `json:",omitempty"`

	Ok  bool  `json:",omitempty"`
	Err error `json:",omitempty"`
}

type UartWhat string

const (
	UartWhatGetDevicePaths UartWhat = "uart.get.device-paths"
	UartWhatOpen           UartWhat = "uart.open"
	UartWhatClose          UartWhat = "uart.close"
	UartWhatSetReadTimeout UartWhat = "uart.set.read.timeout"
	UartWhatRead           UartWhat = "uart.write"
	UartWhatWrite          UartWhat = "uart.read"
)

type Uart struct {

	//
	// Routing part.
	//

	Area Area
	What UartWhat

	//
	// Request part.
	//

	Device string `json:",omitempty"`

	Baudrate int           `json:",omitempty"`
	TimeOut  time.Duration `json:",omitempty"`

	//
	// Response part.
	//

	Paths []string `json:",omitempty"`

	Write []byte `json:",omitempty"`
	Read  []byte `json:",omitempty"`
	Size  int    `json:",omitempty"`
	Xfer  int    `json:",omitempty"`

	Ok  bool  `json:",omitempty"`
	Err error `json:",omitempty"`
}
