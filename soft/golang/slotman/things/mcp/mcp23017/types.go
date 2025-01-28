package mcp23017

import (
	"errors"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
)

type MCP23017 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev *i2c.Device

	handler Handler
}

type Pin byte

type PinDirection byte

const (
	PinDirectionOutput PinDirection = 0
	PinDirectionInput  PinDirection = 1
)

type PinLogic byte

const (
	PinLogicLo PinLogic = 0
	PinLogicHi PinLogic = 1
)

type Control interface {
	SetHandler(handler Handler)

	SetPinDirection(pin Pin, dir PinDirection) (err error)
	GetPinDirection(pin Pin) (dir PinDirection, err error)

	SetPinDirections(directions uint16) (err error)
	GetPinDirections() (directions uint16, err error)

	WritePin(pin Pin, val PinLogic) (err error)
	ReadPin(pin Pin) (val PinLogic, err error)

	WritePins(values uint16) (err error)
	ReadPins() (values uint16, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
}

var (
	ErrInvalidPin   = errors.New("invalid pin")
	ErrInvalidDir   = errors.New("invalid direction")
	ErrInvalidLogic = errors.New("invalid logic")
)
