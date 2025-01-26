package ld2461

import (
	"slotman/drivers/impl/uart"
	"slotman/things"
	"slotman/utils/simple"
)

type LD2461 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string
	BaudRate   int
	IsOpen     bool
	IsStarted  bool

	isProbe bool
	results chan []byte
	uart    *uart.Device
	handler Handler
}

type ReportingFormat byte

//goland:noinspection GoUnusedConst
const (
	ReportingFormatCoordinates ReportingFormat = 1
	ReportingFormatPresence    ReportingFormat = 2
	ReportingFormatBoth        ReportingFormat = 3
)

type ZoneInfo struct {
	Zone int
	Type int

	X0, Y0, X1, Y1, X2, Y2, X3, Y3 float64
}

type Control interface {
	SetHandler(handler Handler)

	FactoryReset() (err error)

	SetBaudRate(baudRate int) (err error)
	GetVersion() (date, version, uid string, err error)
	SetZoneFilter(zi ZoneInfo, active bool) (err error)
	DisableZoneFilter(zone int) (err error)
	ReadZoneFilters() (zi1, zi2, zi3 ZoneInfo, err error)
	SetReportingFormat(format ReportingFormat) (err error)
	GetReportingFormat() (format ReportingFormat, err error)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
	OnHumanTracking(thing things.Thing, xPos, yPos []float64)
}
