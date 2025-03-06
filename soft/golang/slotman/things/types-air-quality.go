package things

import (
	"slotman/utils/simple"
	"time"
)

type AirQuality struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	Percent float64
}

func (aq *AirQuality) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &aq.Uuid
	return
}

func (aq *AirQuality) GetTime() (tp *time.Time) {
	tp = &aq.Date
	return
}

func (aq *AirQuality) SetTime(tp *time.Time) {
	aq.Date = *tp
	return
}

func (aq *AirQuality) GetDay() (want bool) {
	want = true
	return
}

func (aq *AirQuality) GetSub() (sub string) {
	sub = "online/air-quality"
	return
}

func (aq *AirQuality) GetTag() (tag string) {
	tag = "air-quality"
	return
}

func (aq *AirQuality) GetModelInfo() (vendor, model string) {
	vendor = aq.Vendor
	model = aq.Model
	return
}

func (aq *AirQuality) CleanForStorage() {
	aq.What = ""
	aq.Mode = ""
	aq.Uuid = ""
	aq.BoxUuid = ""
	aq.Vendor = ""
	aq.Model = ""
	return
}
