package bmp280

import (
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

var (
	multiOpenLock sync.Mutex
)

type BMP280 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev  *i2c.Device
	handler Handler

	threshold float64
	debug     bool

	digT1 uint16
	digT2 int16
	digT3 int16

	digP1 uint16
	digP2 int16
	digP3 int16
	digP4 int16
	digP5 int16
	digP6 int16
	digP7 int16
	digP8 int16
	digP9 int16

	tFine int
}

type Control interface {
	SetHandler(handler Handler)
	ResetSensor() (err error)
	SetThreshold(threshold float64)
	GetSensorId() (id byte, err error)
	GetStatus() (measuring, imUpdate bool, err error)
	SetMeasureMode(pressOver, tempOver Oversampling) (err error)
	SetIrrFilter(irrFilter IrrFilter) (err error)
	SetPowerMode(pm PowerMode, pi PowerInterval) (err error)
	ReadPressure() (hPa float64, err error)
	ReadTemperature() (celsius float64, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnPressure(thing things.Thing, hPa float64)
	OnTemperature(thing things.Thing, celsius float64)
}
