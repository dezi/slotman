package proxy

import "slotman/utils/simple"

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

	Uuid simple.UUIDHex

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

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}

func (px *Spi) GetUuid() (uuid simple.UUIDHex) {
	uuid = px.Uuid
	return
}

func (px *Spi) SetUuid(uuid simple.UUIDHex) {
	px.Uuid = uuid
	return
}

func (px *Spi) GetArea() (area Area) {
	area = px.Area
	return
}
