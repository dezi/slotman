package things

import (
	"errors"
	"slotman/utils/simple"
)

type ThingType string

const (
	ThingTypePresence     ThingType = "presence"
	ThingTypeFallDetect   ThingType = "fall-detect"
	ThingTypeSleepMonitor ThingType = "sleep-monitor"
	ThingTypeHumanTrack   ThingType = "human-track"
	ThingTypeIllumination ThingType = "illumination"
	ThingTypeRfid         ThingType = "rfid"
	ThingTypePressure     ThingType = "pressure"
	ThingTypeTemperature  ThingType = "temperature"
	ThingTypeRGBColor     ThingType = "rgb-color"
	ThingTypeADConverter  ThingType = "ad-converter"
	ThingTypeIOExpander   ThingType = "io-expander"
	ThingTypeMotorDriver  ThingType = "motor-driver"
	ThingTypeOledDisplay  ThingType = "oled-display"
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
