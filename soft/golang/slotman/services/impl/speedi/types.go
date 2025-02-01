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
	SpeediWhatSpeed SpeediWhat = "speedi.speed"
)

type Speedi struct {

	//
	// Routing part.
	//

	Uuid simple.UUIDHex

	Area proxy.Area
	What SpeediWhat

	//
	// Response part.
	//

	Track    int    `json:",omitempty"`
	RawSpeed uint16 `json:",omitempty"`

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}
