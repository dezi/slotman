package speedi

import (
	"slotman/services/type/proxy"
	"slotman/utils/simple"
)

type SpeedControlCalibration struct {
	MinValue uint16
	MaxValue uint16
}

type PlayerControlCurve struct {
	BrkPercent float64
	MinPercent float64
	MaxPercent float64
}

const (
	AreaSpeedi proxy.Area = "speedi"
)

//goland:noinspection GoNameStartsWithPackageName
type SpeediWhat string

//goland:noinspection GoNameStartsWithPackageName
const (
	SpeediWhatOpen  SpeediWhat = "speedi.open"
	SpeediWhatClose SpeediWhat = "speedi.close"
)

type Speedi struct {

	//
	// Routing part.
	//

	Uuid simple.UUIDHex

	Area proxy.Area
	What SpeediWhat

	//
	// Request part.
	//

	//PinNo uint8 `json:",omitempty"`

	//
	// Response part.
	//

	//State gpio.State `json:",omitempty"`

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}
