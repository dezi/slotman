package sgp40

import (
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
)

type SGP40 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev  *i2c.Device
	handler Handler
	debug   bool
}

type Control interface {
	SetHandler(handler Handler)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	//OnPressure(thing things.Thing, hPa float64)
	//OnTemperature(thing things.Thing, celsius float64)
}
