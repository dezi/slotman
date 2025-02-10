package slotman

import "slotman/utils/simple"

const Version = "v1.0"

type Message struct {
	What string `json:"what"`
}

type Tracks struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Tracks int `json:"tracks"`
}

type Team struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Uuid simple.UUIDHex `json:"uuid"`

	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Car    string `json:"car"`
	CarPic string `json:"carPic"`
}

type Pilot struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Uuid simple.UUIDHex `json:"uuid"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Team      string `json:"team"`
	Car       string `json:"car"`

	MinSpeed int `json:"minSpeed"`
	MaxSpeed int `json:"maxSpeed"`

	ProfilePic string `json:"profilePic"`

	IsMockup bool `json:"-"`
}

type Race struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Title string `json:"title"`

	Tracks int `json:"tracks"`
	Rounds int `json:"rounds"`
}

type Controller struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Controller    int  `json:"controller"`
	IsCalibrating bool `json:"isCalibrating"`

	MinValue int `json:"minValue"`
	MaxValue int `json:"maxValue"`
}

type Setup struct {
	Tracks Tracks
	Race   Race
	Pilots map[simple.UUIDHex]*Pilot
}
