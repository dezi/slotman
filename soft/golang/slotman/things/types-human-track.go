package things

import (
	"slotman/utils/simple"
	"time"
)

type HumanTrack struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	TargetXs []float64
	TargetYs []float64
}

func (ht *HumanTrack) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &ht.Uuid
	return
}

func (ht *HumanTrack) GetTime() (tp *time.Time) {
	tp = &ht.Date
	return
}

func (ht *HumanTrack) SetTime(tp *time.Time) {
	ht.Date = *tp
	return
}

func (ht *HumanTrack) GetDay() (want bool) {
	want = true
	return
}

func (ht *HumanTrack) GetSub() (sub string) {
	sub = "online/human-track"
	return
}

func (ht *HumanTrack) GetTag() (tag string) {
	tag = "human-track"
	return
}

func (ht *HumanTrack) GetModelInfo() (vendor, model string) {
	vendor = ht.Vendor
	model = ht.Model
	return
}

func (ht *HumanTrack) CleanForStorage() {
	ht.What = ""
	ht.Mode = ""
	ht.Uuid = ""
	ht.BoxUuid = ""
	ht.Vendor = ""
	ht.Model = ""
	return
}
