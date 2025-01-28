package ads1115

import (
	"errors"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

type ADS1115 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev   *i2c.Device
	handler  Handler
	readLock sync.Mutex

	rates [4]Rate
	gains [4]Gain

	capMin [4]uint16
	capMax [4]uint16

	epsilon  uint16
	resendMs int64
}

type Control interface {
	SetHandler(handler Handler)

	GetGain() (gain uint16, err error)
	SetGain(gain uint16) (err error)
	GetRate() (rate uint16, err error)
	SetRate(rate uint16) (err error)

	ReadADConversion(input int) (value int16, err error)
}

type Handler interface {
	OnThingOpened(ting things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnADConversion(thing things.Thing, input int, value uint16)
}

var (
	ErrInvalidInput = errors.New("invalid input")
)
