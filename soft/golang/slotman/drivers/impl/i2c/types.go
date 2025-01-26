package i2c

import "os"

type Device struct {
	device string
	addr   uint8
	rc     *os.File
}
