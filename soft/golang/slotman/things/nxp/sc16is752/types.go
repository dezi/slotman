package sc16is752

import (
	"errors"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

type SC15IS752 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev   *i2c.Device
	handler  Handler
	readLock sync.Mutex
}

type Control interface {
	SetHandler(handler Handler)
}

type Handler interface {
	OnThingOpened(ting things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
}

var (
	ErrInvalidInput = errors.New("invalid input")
)
