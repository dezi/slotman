package spi

type Spi interface {
	Open() (err error)
	Close() (err error)

	GetDevice() (device string)
}
