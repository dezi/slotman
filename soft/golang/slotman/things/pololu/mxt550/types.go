package mxt550

import (
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
)

type MXT550 struct {
	Uuid simple.UUIDHex

	Vendor string
	Model  string

	DevicePath string

	IsOpen    bool
	IsStarted bool

	i2cDev *i2c.Device

	handler Handler
	debug   bool
	eeprom  bool

	protocolOptions MotoronProtocolOptions
}

type Control interface {
	SetHandler(handler Handler)

	SetDebug(debug bool)

	EnableCrc() (err error)
	DisableCrc() (err error)
	EnableCrcForCommands() (err error)
	DisableCrcForCommands() (err error)
	EnableCrcForResponses() (err error)
	DisableCrcForResponses() (err error)
	EnableI2cGeneralCall() (err error)
	DisableI2cGeneralCall() (err error)

	SetProtocolOptions(options MotoronProtocolOptions) (err error)
	GetFirmwareVersion() (productId MotoronProductId, firmwareVersion string, err error)

	ReadEeprom(offset, length byte) (buffer []byte, err error)
	ReadEepromDeviceNumber() (number byte, err error)

	WriteEeprom(offset, value byte) (err error)
	WriteEeprom16(offset byte, value uint16) (err error)
	WriteEepromDeviceNumber(number uint16) (err error)
	WriteEepromAlternativeDeviceNumber(number uint16) (err error)
	WriteEepromDisableAlternativeDeviceNumber() (err error)
	WriteEepromCommunicationOptions(options byte) (err error)
	WriteEepromBaudRate(baud uint32) (err error)
	WriteEepromResponseDelay(delay byte) (err error)

	Reinitialize() (err error)
	Reset() (err error)

	GetVariables(motor byte, offset MotoronVar, length byte) (buffer []byte, err error)
	GetVar8(motor byte, offset MotoronVar) (val uint8, err error)
	GetVar16(motor byte, offset MotoronVar) (val uint16, err error)

	GetProtocolOptions() (options MotoronProtocolOptions, err error)

	GetStatusFlags() (flags MotoronStatusFlags, err error)
	GetStatusFlagsStrings(flags MotoronStatusFlags) (flagStrings []string)

	GetProtocolErrorFlag() (flag bool, err error)
	GetCrcErrorFlag() (flag bool, err error)
	GetCommandTimeoutLatchedFlag() (flag bool, err error)
	GetMotorFaultLatchedFlag() (flag bool, err error)
	GetNoPowerLatchedFlag() (flag bool, err error)
	GetUARTErrorFlag() (flag bool, err error)
	GetResetFlag() (flag bool, err error)
	GetCommandTimeoutFlag() (flag bool, err error)
	GetMotorFaultingFlag() (flag bool, err error)
	GetNoPowerFlag() (flag bool, err error)
	GetErrorActiveFlag() (flag bool, err error)
	GetMotorOutputEnabledFlag() (flag bool, err error)
	GetMotorDrivingFlag() (flag bool, err error)

	GetUARTFaults() (val uint8, err error)
	GetVinVoltage() (val uint16, err error)
	GetVinVoltageMv(referenceMv uint16, mType MotoronVinSenseType) (val uint32, err error)
	GetCommandTimeoutMilliseconds() (val uint16, err error)
	GetErrorResponse() (val uint8, err error)
	GetErrorMask() (val uint16, err error)
	GetJumperState() (val uint8, err error)
	GetTargetSpeed(motor byte) (val int16, err error)
	GetTargetBrakeAmount(motor byte) (val uint16, err error)
	GetCurrentSpeed(motor byte) (val int16, err error)
	GetBufferedSpeed(motor byte) (val uint16, err error)
	GetPwmMode(motor byte) (val uint8, err error)
	GetMaxAccelerationForward(motor byte) (val uint16, err error)
	GetMaxAccelerationReverse(motor byte) (val uint16, err error)
	GetMaxDecelerationForward(motor byte) (val uint16, err error)
	GetMaxDecelerationReverse(motor byte) (val uint16, err error)
	GetMaxDecelerationTemporary(motor byte) (val uint16, err error)
	GetStartingSpeedForward(motor byte) (val uint16, err error)
	GetStartingSpeedReverse(motor byte) (val uint16, err error)
	GetDirectionChangeDelayForward(motor byte) (val uint8, err error)
	GetDirectionChangeDelayReverse(motor byte) (val uint8, err error)
	GetCurrentLimit(motor byte) (val uint16, err error)
	GetCurrentSenseReading(motor byte) (mcsr *MotoronCurrentSenseReading, err error)
	GetCurrentSenseRawAndSpeed(motor byte) (mcsr *MotoronCurrentSenseReading, err error)
	GetCurrentSenseProcessedAndSpeed(motor byte) (mcsr *MotoronCurrentSenseReading, err error)
	GetCurrentSenseRaw(motor byte) (senseRaw uint16, err error)
	GetCurrentSenseProcessed(motor byte) (senseProcessed uint16, err error)
	GetCurrentSenseOffset(motor byte) (senseOffset uint8, err error)
	GetCurrentSenseMinimumDivisor(motor byte) (minimumDivisor uint16, err error)
	SetVariable(motor byte, offset MotoronVar, value uint16) (err error)
	SetCommandTimeoutMilliseconds(ms uint16) (err error)
	SetErrorResponse(response uint8) (err error)
	SetErrorMask(mask MotoronStatusFlags) (err error)

	DisableCommandTimeout() (err error)
	ClearUARTFaults(flags uint8) (err error)

	SetPwmMode(motor byte, mode uint8) (err error)
	SetMaxAccelerationForward(motor byte, accel uint16) (err error)
	SetMaxAccelerationReverse(motor byte, accel uint16) (err error)
	SetMaxAcceleration(motor byte, accel uint16) (err error)
	SetMaxDecelerationForward(motor byte, decel uint16) (err error)
	SetMaxDecelerationReverse(motor byte, decel uint16) (err error)
	SetMaxDeceleration(motor byte, decel uint16) (err error)
	SetStartingSpeedForward(motor byte, speed uint16) (err error)
	SetStartingSpeedReverse(motor byte, speed uint16) (err error)
	SetStartingSpeed(motor byte, speed uint16) (err error)
	SetDirectionChangeDelayForward(motor byte, duration uint8) (err error)
	SetDirectionChangeDelayReverse(motor byte, duration uint8) (err error)
	SetDirectionChangeDelay(motor byte, duration uint8) (err error)
	SetCurrentLimit(motor byte, limit uint16) (err error)
	SetCurrentSenseOffset(motor byte, offset uint8) (err error)
	SetCurrentSenseMinimumDivisor(motor byte, speed uint16) (err error)

	CoastNow() (err error)

	ClearMotorFault(flags uint8) (err error)
	ClearMotorFaultUnconditional() (err error)
	ClearLatchedStatusFlags(flags uint16) (err error)
	ClearResetFlag() (err error)

	SetLatchedStatusFlags(flags uint16) (err error)
	SetSpeed(motor byte, speed int16) (err error)
	SetSpeedNow(motor byte, speed int16) (err error)
	SetBufferedSpeed(motor byte, speed int16) (err error)
	SetAllSpeeds(speeds ...int16) (err error)
	SetAllSpeedsNow(speeds ...int16) (err error)
	SetAllBufferedSpeeds(speeds ...int16) (err error)
	SetAllSpeedsUsingBuffers() (err error)
	SetAllSpeedsNowUsingBuffers() (err error)
	SetBraking(motor byte, amount uint16) (err error)
	SetBrakingNow(motor byte, amount uint16) (err error)
	ResetCommandTimeout() (err error)

	CalculateCurrentLimit(
		milliAmps uint32, mType MotoronCurrentSenseType,
		referenceMv uint16, offset uint16) (limit uint16)

	CurrentSenseUnitsMilliamps(
		mType MotoronCurrentSenseType, referenceMv uint16) (milliAmps uint16)
}

type Handler interface {
	OnThingOpened(thing things.Thing)
	OnThingClosed(thing things.Thing)
	OnThingStarted(thing things.Thing)
	OnThingStopped(thing things.Thing)
}
