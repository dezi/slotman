package proxy

import "slotman/utils/simple"

type Area string

const (
	AreaI2c  Area = "i2c"
	AreaSpi  Area = "spi"
	AreaUart Area = "uart"
	AreaGpio Area = "gpio"
)

type Message interface {
	GetUuid() (uuid simple.UUIDHex)
	SetUuid(uuid simple.UUIDHex)
	GetArea() (area Area)
}

type Subscriber interface {
	OnMessageFromClient(reqBytes []byte) (resBytes []byte, err error)
	OnMessageFromServer(resBytes []byte)
}
