package tacho

import (
	"slotman/services/type/proxy"
	"slotman/utils/simple"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type TachoState struct {
	active bool
	dirty  bool
	time   time.Time
}

//goland:noinspection GoNameStartsWithPackageName
type TachoRead struct {
	pinStates uint16
	readTime  time.Time
}

type TrackState struct {
	Round       int
	RoundMillis int
	RoundTs     time.Time

	SpeedKmh float64
	SpeedTS  time.Time

	IsAtStart bool
}

const (
	AreaTacho proxy.Area = "tacho"
)

//goland:noinspection GoNameStartsWithPackageName
type TachoWhat string

//goland:noinspection GoNameStartsWithPackageName
const (
	TachoWhatSpeed TachoWhat = "tacho.speed"
)

type Tacho struct {

	//
	// Routing part.
	//

	Uuid simple.UUIDHex

	Area proxy.Area
	What TachoWhat

	//
	// Response part.
	//

	Pin    byte      `json:",omitempty"`
	Active bool      `json:",omitempty"`
	Time   time.Time `json:",omitempty"`

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}
