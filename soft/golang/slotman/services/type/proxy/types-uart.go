package proxy

import (
	"slotman/utils/simple"
	"time"
)

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

	Uuid simple.UUIDHex

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

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}

func (px *Uart) GetUuid() (uuid simple.UUIDHex) {
	uuid = px.Uuid
	return
}

func (px *Uart) SetUuid(uuid simple.UUIDHex) {
	px.Uuid = uuid
	return
}

func (px *Uart) GetArea() (area Area) {
	area = px.Area
	return
}
