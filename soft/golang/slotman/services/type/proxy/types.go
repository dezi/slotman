package proxy

import "slotman/utils/simple"

type Area string

const (
	AreaGpio Area = "gpio"
	AreaI2c  Area = "i2c"
	AreaSpi  Area = "spi"
	AreaUart Area = "uart"
)

type Message struct {
	Uuid simple.UUIDHex
	Area Area
}

type MessageIface interface {
	GetUuid() (uuid simple.UUIDHex)
	SetUuid(uuid simple.UUIDHex)
	GetArea() (area Area)
}
