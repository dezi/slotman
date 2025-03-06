package ld6001a

import (
	"slotman/drivers/iface/uart"
	"slotman/things"
	"slotman/utils/simple"
)

type LD6001a struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string
	BaudRate   int
	IsOpen     bool
	IsStarted  bool

	isProbe bool
	results chan string
	uart    uart.Uart
	handler Handler
}

type Control interface {
	SetHandler(handler Handler)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnHumanPresence(thing things.Thing, people int)
	OnHumanTracking3D(thing things.Thing, xPos, yPos, zPos, xV, yV, zV []float64)
}
