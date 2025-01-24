package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

type State uint8

const (
	Low State = iota
	High
)

type Pin struct {
	pin     rpio.Pin
	pinNo   uint8
	on      time.Duration
	off     time.Duration
	loops   int
	started bool
}
