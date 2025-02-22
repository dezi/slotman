package sgp40

import (
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

var (
	multiOpenLock sync.Mutex
)

type SGP40 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev  *i2c.Device
	lock    sync.Mutex
	handler Handler
	debug   bool

	humidity    int
	temperature int
}

type Control interface {
	SetHandler(handler Handler)

	DoSelfTest() (ok bool, err error)
	ReadSerial() (serial string, err error)

	SetHumidity(percent int) (err error)
	SetTemperature(celsius int) (err error)

	MeasureRawSignal() (signal, rawSignal int, err error)
	MeasureAirQuality() (percent float64, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnAirQuality(thing things.Thing, percent float64)
}
