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

	i2cDev      *i2c.Device
	handler     Handler
	accessLock  sync.Mutex
	crystalFreq int

	pollSleep    [2]int
	readTimeout  [2]int
	writeTimeout [2]int
}

type Control interface {
	SetHandler(handler Handler)

	Ping() (err error)

	SetFifoEnable(channel byte, enable bool) (err error)
	SetCrystalFreq(crystalFreq int) (err error)
	SetBaudrate(channel byte, baudrate int) (err error)
	SetLine(channel byte, dataBits, parity, stopBits byte) (err error)
	SetPollInterval(channel byte, millis int) (err error)
	SetReadTimeout(channel byte, millis int) (err error)
	SetWriteTimeout(channel byte, millis int) (err error)

	WriteUartByte(channel, value byte) (err error)
	WriteUartBytes(channel byte, data []byte) (xfer int, err error)

	ReadUartByte(channel byte) (value byte, err error)
	ReadUartBytes(channel byte, size int) (xfer int, data []byte, err error)

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
	ErrReadTimeout    = errors.New("read timeout")
	ErrWriteTimeout   = errors.New("write timeout")
	ErrDeviceClosed   = errors.New("device closed")
)
