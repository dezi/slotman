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

type Control interface {
}

type Handler interface {
	//OnSensorOpened(sensor sensors.Sensor)
	//OnSensorClosed(sensor sensors.Sensor)
	//OnSensorStarted(sensor sensors.Sensor)
	//OnSensorStopped(sensor sensors.Sensor)
}
