package proxy

import "slotman/utils/simple"

type I2cWhat string

const (
	I2cWhatGetDevicePaths I2cWhat = "i2c.get.device-paths"
	I2cWhatOpen           I2cWhat = "i2c.open"
	I2cWhatClose          I2cWhat = "i2c.close"
	I2cWhatTransLock      I2cWhat = "i2c.trans.lock"
	I2cWhatTransUnlock    I2cWhat = "i2c.trans.unlock"
	I2cWhatWrite          I2cWhat = "i2c.write"
	I2cWhatRead           I2cWhat = "i2c.read"
	I2cWhatReadUart       I2cWhat = "i2c.read.uart"
)

type I2c struct {

	//
	// Routing part.
	//

	Uuid simple.UUIDHex

	Area Area
	What I2cWhat

	//
	// Request part.
	//

	Device string `json:",omitempty"`
	Addr   uint8  `json:",omitempty"`

	//
	// Request for fast uart read.
	//

	Channel byte `json:",omitempty"`
	TimeOut int  `json:",omitempty"`

	//
	// Response part.
	//

	Paths []string `json:",omitempty"`
	Write []byte   `json:",omitempty"`
	Read  []byte   `json:",omitempty"`
	Size  int      `json:",omitempty"`
	Xfer  int      `json:",omitempty"`

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}

func (px *I2c) GetUuid() (uuid simple.UUIDHex) {
	uuid = px.Uuid
	return
}

func (px *I2c) SetUuid(uuid simple.UUIDHex) {
	px.Uuid = uuid
	return
}

func (px *I2c) GetArea() (area Area) {
	area = px.Area
	return
}
