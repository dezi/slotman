package spi

type Spi interface {
	Open() (err error)
	Close() (err error)

	GetDevice() (device string)

	SetMode(mode uint8) (err error)
	SetBitsPerWord(bpw uint8) (err error)
	SetSpeed(speed uint32) (err error)

	Send(request []byte) (response []byte, err error)
}
