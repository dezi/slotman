package turner

import (
	"slotman/services/type/proxy"
	"slotman/utils/simple"
)

const (
	AreaTurner proxy.Area = "turner"
)

//goland:noinspection GoNameStartsWithPackageName
type TurnerWhat string

//goland:noinspection GoNameStartsWithPackageName
const (
	TurnerWhatBlipFull TurnerWhat = "turner.blip.full"
)

type Turner struct {

	//
	// Routing part.
	//

	Uuid simple.UUIDHex

	Area proxy.Area
	What TurnerWhat

	//
	// Request part.
	//

	BlipImage []byte `json:",omitempty"`

	//
	// Response part.
	//

	Ok  bool   `json:",omitempty"`
	Err string `json:",omitempty"`

	NE error `json:"-"`
}

func (tr *Turner) GetUuid() (uuid simple.UUIDHex) {
	uuid = tr.Uuid
	return
}

func (tr *Turner) SetUuid(uuid simple.UUIDHex) {
	tr.Uuid = uuid
	return
}

func (tr *Turner) GetArea() (area proxy.Area) {
	area = tr.Area
	return
}
