package uart

import "time"

type Uart interface {
	Open() (err error)
	Close() (err error)

	GetDevice() (device string)
	GetBaudrate() (baudrate int)

	SetBaudrate(baudrate int) (err error)
	SetReadTimeout(timeout time.Duration) (err error)

	Write(data []byte) (xfer int, err error)
	Read(data []byte) (xfer int, err error)
}
