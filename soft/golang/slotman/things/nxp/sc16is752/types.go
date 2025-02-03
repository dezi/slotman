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

	pollSleep [2]int
}

type Control interface {
	SetHandler(handler Handler)

	Ping() (err error)

	EnableFifo(channel byte, enable bool) (err error)

	ReadRegister(register, channel byte) (value byte, err error)
	WriteRegister(register, channel, value byte) (err error)
}

type Handler interface {
	OnThingOpened(ting things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)

	OnUartDataReceived(thing things.Thing, channel byte, data []byte)
}

var (
	ErrInvalidChannel = errors.New("invalid channel")
	ErrInvalidPing    = errors.New("invalid ping")
)
