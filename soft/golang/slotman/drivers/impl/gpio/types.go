package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Pin struct {
	pin   rpio.Pin
	pinNo uint8
}
