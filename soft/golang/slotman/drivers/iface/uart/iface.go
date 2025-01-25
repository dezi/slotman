package uart

import "time"

type Uart interface {
	Open() (err error)
	Close() (err error)

	GetDevice() (device string)

	SetReadTimeout(timeout time.Duration) (err error)

	Read(data []byte) (xfer int, err error)
	Write(data []byte) (xfer int, err error)
}
