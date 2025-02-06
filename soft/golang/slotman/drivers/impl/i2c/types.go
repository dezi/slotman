package i2c

import (
	"os"
	"sync"
)

type Device struct {
	device string
	addr   uint8
	rc     *os.File
	lock   sync.Mutex
}
