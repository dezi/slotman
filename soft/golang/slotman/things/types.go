package things

import (
	"errors"
	"slotman/utils/simple"
)

type ThingType string

const (
	ThingTypePresence      ThingType = "tt-presence"
	ThingTypeFallDetect    ThingType = "tt-fall-detect"
	ThingTypeSleepMonitor  ThingType = "tt-sleep-monitor"
	ThingTypeHumanTrack    ThingType = "tt-human-track"
	ThingTypeIllumination  ThingType = "tt-illumination"
	ThingTypeRfid          ThingType = "tt-rfid"
	ThingTypePressure      ThingType = "tt-pressure"
	ThingTypeHumidity      ThingType = "tt-humidity"
	ThingTypeTemperature   ThingType = "tt-temperature"
	ThingTypeRGBColor      ThingType = "tt-rgb-color"
	ThingTypeADConverter   ThingType = "tt-ad-converter"
	ThingTypeIOExpander    ThingType = "tt-io-expander"
	ThingTypeMotorDriver   ThingType = "tt-motor-driver"
	ThingTypeColorDisplay  ThingType = "tt-color-display"
	ThingTypeUartConverter ThingType = "tt-uart-converter"
)

type Thing interface {
	GetUuid() (uuid simple.UUIDHex)
	GetModelInfo() (vendor, model, short string)
	GetThingTypes() (ThingTypes []ThingType)
	GetThingDevicePath() (devicePath string)
	GetThingAddress() (address int)

	Open() (err error)
	Close() (err error)
	Start() (err error)
	Stop() (err error)
}

type Handler interface {
	OnThingOpened(thing Thing)
	OnThingClosed(thing Thing)
	OnThingStarted(thing Thing)
	OnThingStopped(thing Thing)
}

var (
	ErrThingNotOpen        = errors.New("thing not open")
	ErrThingNotStarted     = errors.New("thing not started")
	ErrUnsupportedBaudRate = errors.New("unsupported baud rate")
)
