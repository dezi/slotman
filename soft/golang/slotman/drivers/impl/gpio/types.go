package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Pin struct {
	PinNo uint8
	pin   rpio.Pin
}
