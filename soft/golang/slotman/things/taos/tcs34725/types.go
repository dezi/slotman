package tcs34725

import (
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

var (
	multiOpenLock sync.Mutex
)

type TCS34725 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev    *i2c.Device
	handler   Handler
	threshold float64
}

type Control interface {
	SetHandler(handler Handler)

	InitThing(gain Gain, it IntegrationTime) (err error)
	SetEnabled(enabled bool) (err error)
	SetThreshold(threshold float64)
	GetGain() (gain Gain, err error)
	SetGain(gain Gain) (err error)
	GetIntegrationTime() (it IntegrationTime, err error)
	SetIntegrationTime(it IntegrationTime) (err error)
	ReadRgbColor() (r, g, b, lux int, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnRGBColor(thing things.Thing, r, g, b, lux int)
}
