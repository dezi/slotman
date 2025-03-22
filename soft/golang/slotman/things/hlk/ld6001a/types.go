package ld6001a

import (
	"errors"
	"slotman/drivers/iface/uart"
	"slotman/things"
	"slotman/utils/simple"
	"sync"
)

type LD6001a struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string
	BaudRate   int
	IsOpen     bool
	IsStarted  bool

	Params ReadParams

	isProbe   bool
	results   chan string
	uart      uart.Uart
	handler   Handler
	loopGroup sync.WaitGroup
}

type Control interface {
	SetHandler(handler Handler)

	StartWorking() (err error)
	StopWorking() (err error)

	Reset() (err error)
	Restore() (err error)

	ReadParams() (err error)

	SetBaudRate(baudRate int) (err error)
	SetSensitivity(sensitivity int) (err error)
	SetRange(xrange int) (err error)
	SetHeight(height int) (err error)
	SetProtocol(mode int) (err error)
	SetRanges(xPosi, xNega, yPosi, yNega int) (err error)
	SetMoving(time int) (err error)
	SetStatic(time int) (err error)
	SetExit(time int) (err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnHumanTracking(thing things.Thing, xPos, yPos []float64)
	OnHumanTracking3D(thing things.Thing, ids []int, xps, yps, zps, xvs, yvs, zvs []float64)
}

// {
// "PeopleCntSoftVerison":"NOP_1.07",
// "RangeRes":0.055664,
// "VelRes":0.111289,
// "TIME":100,
// "PROG":2,
// "Range":450,
// "Sen":4,
// "Heart_Time":60,
// "Debug":1,
// "detectionHeight":300,
// "XboundaryN":-450,
// "XboundaryP":450,
// "YboundaryN":-450,
// "YboundaryP":450,
// "Moving target": 11.00,
// "Static target": 10.00,
// "Target exit":  0.50,
// }

type ReadParams struct {
	SoftwareVersion    string  `json:"PeopleCntSoftVerison"`
	RangeRes           float64 `json:"RangeRes"`
	VelRes             float64 `json:"VelRes"`
	Time               int     `json:"TIME"`
	Prog               int     `json:"PROG"`
	Range              int     `json:"Range"`
	LongSensitivity    int     `json:"Sen"`
	HeartBeatIntervall int     `json:"Heart_Time"`
	Debug              int     `json:"Debug"`
	DetectionHeight    int     `json:"detectionHeight"`
	XNega              int     `json:"XboundaryN"`
	XPosi              int     `json:"XboundaryP"`
	YNega              int     `json:"YboundaryN"`
	YPosi              int     `json:"YboundaryP"`
	MovingTargetDisapp float64 `json:"Moving target"`
	StaticTargetDisapp float64 `json:"Static target"`
	TargetExit         float64 `json:"Target exit"`
}

var (
	ErrSerialTimeout = errors.New("serial timeout")
)
