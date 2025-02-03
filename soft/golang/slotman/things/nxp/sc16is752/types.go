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

	Ping() (err error)
}

type Handler interface {
	OnThingOpened(ting things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
}

var (
	ErrInvalidChannel = errors.New("invalid channel")
	ErrInvalidPing    = errors.New("invalid ping")
)
