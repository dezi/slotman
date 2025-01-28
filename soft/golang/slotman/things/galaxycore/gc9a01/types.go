package gc9a01

import (
	"image"
	"slotman/drivers/iface/gpio"
	"slotman/drivers/impl/spi"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

type GC9A01 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	dcPinNo byte
	dcPin   gpio.Gpio
	spiDev  *spi.Device

	blipLast []byte
	blipLock sync.Mutex

	handler Handler
	debug   bool
}

type Frame struct {
	X0 uint16
	Y0 uint16
	X1 uint16
	Y1 uint16
}

type Control interface {
	SetHandler(handler Handler)

	Initialize() (err error)

	SetFrame(frame Frame) (err error)

	BlipFullImage(img image.Image) (err error)
	BlipFullImageRaw(image []byte) (err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
}
