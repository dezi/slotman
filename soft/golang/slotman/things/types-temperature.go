package things

import (
	"slotman/utils/simple"
	"time"
)

type Temperature struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	Celsius float64
}

func (te *Temperature) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &te.Uuid
	return
}

func (te *Temperature) GetTime() (tp *time.Time) {
	tp = &te.Date
	return
}

func (te *Temperature) SetTime(tp *time.Time) {
	te.Date = *tp
	return
}

func (te *Temperature) GetDay() (want bool) {
	want = true
	return
}

func (te *Temperature) GetSub() (sub string) {
	sub = "online/temperature"
	return
}

func (te *Temperature) GetTag() (tag string) {
	tag = "temperature"
	return
}

func (te *Temperature) GetModelInfo() (vendor, model string) {
	vendor = te.Vendor
	model = te.Model
	return
}

func (te *Temperature) CleanForStorage() {
	te.What = ""
	te.Mode = ""
	te.Uuid = ""
	te.BoxUuid = ""
	te.Vendor = ""
	te.Model = ""
	return
}
