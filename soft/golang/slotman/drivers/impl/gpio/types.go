package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type State uint8

const (
	Low State = iota
	High
)

type Pin struct {
	Pin   rpio.Pin
	PinNo uint8
}
