package gps6mv2

import (
	"slotman/drivers/iface/uart"
	"slotman/things"
	"slotman/utils/simple"
)

type GPS6MV2 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string
	BaudRate   int
	IsOpen     bool
	IsStarted  bool

	Latitude  float64
	Longitude float64
	Elevation float64

	isProbe bool
	results chan string
	uart    uart.Uart
	handler Handler
}

type Control interface {
	SetHandler(handler Handler)
	GetCurrentPosition() (latitude, longitude, elevation float64, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnGPSPosition(thing things.Thing, lat, lon, ele float64)
}
