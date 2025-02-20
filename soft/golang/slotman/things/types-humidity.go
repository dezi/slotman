package things

import (
	"slotman/utils/simple"
	"time"
)

type Humidity struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	Percent float64
}

func (hu *Humidity) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &hu.Uuid
	return
}

func (hu *Humidity) GetTime() (tp *time.Time) {
	tp = &hu.Date
	return
}

func (hu *Humidity) SetTime(tp *time.Time) {
	hu.Date = *tp
	return
}

func (hu *Humidity) GetDay() (want bool) {
	want = true
	return
}

func (hu *Humidity) GetSub() (sub string) {
	sub = "online/humidity"
	return
}

func (hu *Humidity) GetTag() (tag string) {
	tag = "humidity"
	return
}

func (hu *Humidity) GetModelInfo() (vendor, model string) {
	vendor = hu.Vendor
	model = hu.Model
	return
}

func (hu *Humidity) CleanForStorage() {
	hu.What = ""
	hu.Mode = ""
	hu.Uuid = ""
	hu.BoxUuid = ""
	hu.Vendor = ""
	hu.Model = ""
	return
}
