package ld2451

import (
	"slotman/drivers/iface/uart"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

type LD2451 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string
	BaudRate   int
	IsOpen     bool
	IsStarted  bool

	isProbe bool
	results chan []byte
	uart    uart.Uart
	lock    sync.Mutex

	handler Handler
}

type Control interface {
	SetHandler(handler Handler)

	EnableConfigurations() (protocol int, err error)
	EndConfiguration() (err error)

	FactoryReset() (err error)
	RestartModule() (err error)
	SetBaudRate(baudRate int) (err error)

	GetVersion() (fType, major, minor int, err error)

	GetDetectionParams() (
		maxDist, minSpeed, delay byte,
		mode DetectionMode, err error)

	SetDetectionParams(
		maxDist, minSpeed, delay byte,
		mode DetectionMode) (err error)

	GetSensitivityParams() (trigger, noise byte, err error)
	SetSensitivityParams(trigger, noise byte) (err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnHumanTracking(thing things.Thing, xPos, yPos []float64)
}
