package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Pin struct {
	Pin   rpio.Pin
	PinNo uint8
}
