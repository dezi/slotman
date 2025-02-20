package things

import (
	"slotman/utils/simple"
	"time"
)

type GpsPosition struct {
	What ThingType `json:",omitempty"`
	Mode string    `json:",omitempty"`

	Uuid    simple.UUIDHex `json:",omitempty"`
	Date    time.Time      `json:",omitempty"`
	BoxUuid simple.UUIDHex `json:",omitempty"`

	Vendor string `json:",omitempty"`
	Model  string `json:",omitempty"`

	Lat float64
	Lon float64
	Ele float64
}

func (gps *GpsPosition) GetUuid() (uuid *simple.UUIDHex) {
	uuid = &gps.Uuid
	return
}

func (gps *GpsPosition) GetTime() (tp *time.Time) {
	tp = &gps.Date
	return
}

func (gps *GpsPosition) SetTime(tp *time.Time) {
	gps.Date = *tp
	return
}

func (gps *GpsPosition) GetDay() (want bool) {
	want = true
	return
}

func (gps *GpsPosition) GetSub() (sub string) {
	sub = "online/gps-position"
	return
}

func (gps *GpsPosition) GetTag() (tag string) {
	tag = "gps-position"
	return
}

func (gps *GpsPosition) GetModelInfo() (vendor, model string) {
	vendor = gps.Vendor
	model = gps.Model
	return
}

func (gps *GpsPosition) CleanForStorage() {
	gps.What = ""
	gps.Mode = ""
	gps.Uuid = ""
	gps.BoxUuid = ""
	gps.Vendor = ""
	gps.Model = ""
	return
}
