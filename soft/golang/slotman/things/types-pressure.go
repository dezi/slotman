package things

import (
	"slotman/utils/simple"
	"time"
)

type Pressure struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	HPa float64
}

func (pr *Pressure) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &pr.Uuid
	return
}

func (pr *Pressure) GetTime() (tp *time.Time) {
	tp = &pr.Date
	return
}

func (pr *Pressure) SetTime(tp *time.Time) {
	pr.Date = *tp
	return
}

func (pr *Pressure) GetDay() (want bool) {
	want = true
	return
}

func (pr *Pressure) GetSub() (sub string) {
	sub = "online/pressure"
	return
}

func (pr *Pressure) GetTag() (tag string) {
	tag = "pressure"
	return
}

func (pr *Pressure) GetModelInfo() (vendor, model string) {
	vendor = pr.Vendor
	model = pr.Model
	return
}

func (pr *Pressure) CleanForStorage() {
	pr.What = ""
	pr.Mode = ""
	pr.Uuid = ""
	pr.BoxUuid = ""
	pr.Vendor = ""
	pr.Model = ""
	return
}
