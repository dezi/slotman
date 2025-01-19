package gc9a01

import (
	"slotman/drivers/gpio"
	"slotman/drivers/spi"
)

type GC9A01 struct {
	//Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	dcPinNo byte
	dcPin   *gpio.Pin
	spi     *spi.Device
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
}

type Handler interface {
	//OnSensorOpened(sensor sensors.Sensor)
	//OnSensorClosed(sensor sensors.Sensor)
	//OnSensorStarted(sensor sensors.Sensor)
	//OnSensorStopped(sensor sensors.Sensor)
}
