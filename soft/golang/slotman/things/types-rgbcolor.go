package things

import (
	"slotman/utils/simple"
	"time"
)

type RGBColor struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	R int
	G int
	B int

	Lux int
}

func (rgb *RGBColor) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &rgb.Uuid
	return
}

func (rgb *RGBColor) GetTime() (tp *time.Time) {
	tp = &rgb.Date
	return
}

func (rgb *RGBColor) SetTime(tp *time.Time) {
	rgb.Date = *tp
	return
}

func (rgb *RGBColor) GetDay() (want bool) {
	want = true
	return
}

func (rgb *RGBColor) GetSub() (sub string) {
	sub = "online/rgb-color"
	return
}

func (rgb *RGBColor) GetTag() (tag string) {
	tag = "rgb-color"
	return
}

func (rgb *RGBColor) GetModelInfo() (vendor, model string) {
	vendor = rgb.Vendor
	model = rgb.Model
	return
}

func (rgb *RGBColor) CleanForStorage() {
	rgb.What = ""
	rgb.Mode = ""
	rgb.Uuid = ""
	rgb.BoxUuid = ""
	rgb.Vendor = ""
	rgb.Model = ""
	return
}
