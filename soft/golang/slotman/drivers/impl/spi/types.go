package spi

import "os"

type Device struct {
	Path string

	file  *os.File
	mode  uint8
	bpw   uint8
	speed uint32
	delay uint16
}
