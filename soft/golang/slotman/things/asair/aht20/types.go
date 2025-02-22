package aht20

import (
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
)

type AHT20 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev *i2c.Device

	handler   Handler
	threshold float64

	debug bool
}

type Control interface {
	SetHandler(handler Handler)

	Init() (err error)
	Reset() (err error)

	SetThreshold(threshold float64)
	ReadMeasurement() (humidity, celsius float64, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnTemperature(thing things.Thing, celsius float64)
	OnHumidity(thing things.Thing, percent float64)
}
