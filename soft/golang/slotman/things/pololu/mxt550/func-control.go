package mxt550

import (
	"fmt"
	"slotman/utils/log"
	"time"
)

func (se *MXT550) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *MXT550) SetDebug(debug bool) {
	se.debug = debug
}

// EnableCrc
// Enables CRC for commands and responses.
// see also: SetProtocolOptions(),
func (se *MXT550) EnableCrc() (err error) {
	options := se.protocolOptions
	options |= 1 << MotoronProtocolOptionCrcForCommands
	options |= 1 << MotoronProtocolOptionCrcForResponses
	err = se.SetProtocolOptions(options)
	return
}

// DisableCrc
// Disables CRC for commands and responses.
// see also: SetProtocolOptions(),
func (se *MXT550) DisableCrc() (err error) {
	options := se.protocolOptions
	options ^= 1 << MotoronProtocolOptionCrcForCommands
	options ^= 1 << MotoronProtocolOptionCrcForResponses
	err = se.SetProtocolOptions(options)
	return
}

// EnableCrcForCommands
// Enables CRC for commands.
// see also: SetProtocolOptions(),
func (se *MXT550) EnableCrcForCommands() (err error) {
	options := se.protocolOptions
	options |= 1 << MotoronProtocolOptionCrcForCommands
	err = se.SetProtocolOptions(options)
	return
}

// DisableCrcForCommands
// Disables CRC for commands.
// see also: SetProtocolOptions(),
func (se *MXT550) DisableCrcForCommands() (err error) {
	options := se.protocolOptions
	options ^= 1 << MotoronProtocolOptionCrcForCommands
	err = se.SetProtocolOptions(options)
	return
}

// EnableCrcForResponses
// Enables CRC for responses.
// see also: SetProtocolOptions(),
func (se *MXT550) EnableCrcForResponses() (err error) {
	options := se.protocolOptions
	options |= 1 << MotoronProtocolOptionCrcForResponses
	err = se.SetProtocolOptions(options)
	return
}

// DisableCrcForResponses
// Disables CRC for responses.
// see also: SetProtocolOptions(),
func (se *MXT550) DisableCrcForResponses() (err error) {
	options := se.protocolOptions
	options ^= 1 << MotoronProtocolOptionCrcForResponses
	err = se.SetProtocolOptions(options)
	return
}

// EnableI2cGeneralCall
// Enables the I2C general call address.
// see also: SetProtocolOptions(),
func (se *MXT550) EnableI2cGeneralCall() (err error) {
	options := se.protocolOptions
	options |= 1 << MotoronProtocolOptionI2cGeneralCall
	err = se.SetProtocolOptions(options)
	return
}

// DisableI2cGeneralCall
// Disables the I2C general call address.
// see also: SetProtocolOptions(),
func (se *MXT550) DisableI2cGeneralCall() (err error) {
	options := se.protocolOptions
	options ^= 1 << MotoronProtocolOptionI2cGeneralCall
	err = se.SetProtocolOptions(options)
	return
}

// SetProtocolOptions
//
// Sends the "Set protocol options" command to the device to specify options
// related to how the device processes commands and sends responses.
// The options are also saved in this object and are used later
// when sending other commands or reading responses.
//
// When CRC for commands is enabled, this library generates the CRC
// byte and appends it to the end of each command it sends.  The Motoron
// checks it to help ensure the command was received correctly.
//
// When CRC for responses is enabled, this library reads the CRC byte sent
// by the Motoron in its responses and makes sure it is correct.  If the
// response CRC byte is incorrect, getLastError() will return a non-zero
// error code after the command has been run.
//
// When the I2C general call address is enabled, the Motoron receives
// commands sent to address 0 in addition to its usual I2C address.
// The general call address is write-only; reading bytes from it is not
// supported.
//
// By default, (in this library and the Motoron itself) CRC for commands and
// responses is enabled, and the I2C general call address is enabled.
//
// This method always sends its command with a CRC byte, so it will work
// even if CRC was previously disabled but has been re-enabled on the device
// (e.g. due to a reset).
//
// The options argument should be 0 or a combination of the following
// expressions made using the bitwise or operator (|):
// - (1 << MotoronProtocolOptionCrcForCommands)
// - (1 << MotoronProtocolOptionCrcForResponses)
// - (1 << MotoronProtocolOptionI2cGeneralCall)
//
// For more information, see the "Set protocol options"
// command in the Motoron user's guide.
//
// see also:
// EnableCrc(), DisableCrc(),
// EnableCrcForCommands(), DisableCrcForCommands(),
// EnableCrcForResponses(), DisableCrcForResponses(),
// EnableI2cGeneralCall(), DisableI2cGeneralCall()
func (se *MXT550) SetProtocolOptions(options MotoronProtocolOptions) (err error) {

	se.protocolOptions = options

	cmd := []byte{
		byte(MotoronCmdSetProtocolOptions),
		byte(options) & 0x7f,
		^byte(options) & 0x7f,
	}

	err = se.sendCommandCrc(cmd)
	return
}

// GetFirmwareVersion
//
// Sends the "Get firmware version" command to get the device's firmware
// product ID and firmware version numbers.
//
// For more information, see the "Get firmware version"
// command in the Motoron user's guide.
func (se *MXT550) GetFirmwareVersion() (productId MotoronProductId, firmwareVersion string, err error) {

	response, err := se.sendCommandAndReadResponse([]byte{byte(MotoronCmdGetFirmwareVersion)}, 4)
	if err != nil {
		return
	}

	productId = MotoronProductId(response[0]) | MotoronProductId(response[1])<<8
	firmwareVersion = fmt.Sprintf("%d.%d", response[2], response[3])

	return
}

// ReadEeprom
//
// Reads the specified bytes from the Motoron's EEPROM memory.
//
// For more information, see the "Read EEPROM" command in the
// Motoron user's guide.
func (se *MXT550) ReadEeprom(offset, length byte) (buffer []byte, err error) {

	cmd := []byte{
		byte(MotoronCmdReadEeprom),
		offset & 0x7F,
		length & 0x7F,
	}

	buffer, err = se.sendCommandAndReadResponse(cmd, length)
	return
}

// ReadEepromDeviceNumber
//
// Reads the EEPROM device number from the device.
// This is the I2C address that the device uses if it detects that JMP1
// is shorted to GND when it starts up.  It is stored in non-volatile
// EEPROM memory.
func (se *MXT550) ReadEepromDeviceNumber() (number byte, err error) {
	buffer, err := se.ReadEeprom(byte(MotoronSettingDeviceNumber), 1)
	if err != nil {
		return
	}
	number = buffer[0]
	return
}

// WriteEeprom
//
// Writes a value to one byte in the Motoron's EEPROM memory.
//
// This command only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
//
// For more information, see the "Write EEPROM" command in the
// Motoron user's guide.
func (se *MXT550) WriteEeprom(offset, value byte) (err error) {

	var cmd [7]byte
	cmd[0] = byte(MotoronCmdWriteEeprom)
	cmd[1] = offset & 0x7F
	cmd[2] = value & 0x7F
	cmd[3] = value >> 7 & 1
	cmd[4] = cmd[1] ^ 0x7F
	cmd[5] = cmd[2] ^ 0x7F
	cmd[6] = cmd[3] ^ 0x7F

	err = se.sendCommand(cmd[:])
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 10)
	return
}

// WriteEeprom16
//
// Writes a 2-byte value in the Motoron's EEPROM memory.
//
// This command only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
func (se *MXT550) WriteEeprom16(offset byte, value uint16) (err error) {

	err = se.WriteEeprom(offset+0, byte(value))
	if err != nil {
		return
	}

	err = se.WriteEeprom(offset+1, byte(value>>8))

	return
}

// WriteEepromDeviceNumber
//
// Writes to the device number stored in EEPROM, changing it to the
// specified value.
//
// This command only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
func (se *MXT550) WriteEepromDeviceNumber(number byte) (err error) {

	err = se.WriteEeprom(byte(MotoronSettingDeviceNumber), number&0x7f)
	if err != nil {
		return
	}

	return
}

// WriteEepromAlternativeDeviceNumber
//
// Writes to the alternative device number stored in EEPROM, changing it to
// the specified value.
//
// This function is only useful for Motorons with a UART serial interface,
// and only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
//
// See also: WriteEepromDisableAlternativeDeviceNumber()
func (se *MXT550) WriteEepromAlternativeDeviceNumber(number uint16) (err error) {

	err = se.WriteEeprom(byte(MotoronSettingAlternativeDeviceNumber), byte((number&0x7F)|0x80))
	if err != nil {
		return
	}

	err = se.WriteEeprom(byte(MotoronSettingAlternativeDeviceNumber)+1, byte(number>>7&0x7F))
	return
}

// WriteEepromDisableAlternativeDeviceNumber
//
// Writes to EEPROM to disable the alternative device number.
//
// This function is only useful for Motorons with a UART serial interface,
// and only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
//
// See also: WriteEepromAlternativeDeviceNumber()
func (se *MXT550) WriteEepromDisableAlternativeDeviceNumber() (err error) {

	err = se.WriteEeprom(byte(MotoronSettingAlternativeDeviceNumber), 0)
	if err != nil {
		return
	}

	err = se.WriteEeprom(byte(MotoronSettingAlternativeDeviceNumber)+1, 0)
	return
}

// WriteEepromCommunicationOptions
//
// Writes to the communication options byte stored in EEPROM, changing it to
// the specified value.
//
// The bits in this byte are defined by the MOTORON_COMMUNICATION_OPTION_*
// macros.
//
// This function is only useful for Motorons with a UART serial interface,
// and only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
func (se *MXT550) WriteEepromCommunicationOptions(options byte) (err error) {
	err = se.WriteEeprom(byte(MotoronSettingCommunicationOptions), options)
	return
}

// WriteEepromBaudRate
//
// Writes to the baud rate stored in EEPROM, changing it to the
// specified value.
//
// This function is only useful for Motorons with a UART serial interface,
// and only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
func (se *MXT550) WriteEepromBaudRate(baud uint32) (err error) {

	if baud < MotoronMinBaudRate {
		baud = MotoronMinBaudRate
	}

	if baud > MotoronMaxBaudRate {
		baud = MotoronMaxBaudRate
	}

	err = se.WriteEeprom16(byte(MotoronSettingBaudDivider), uint16((16000000+(baud>>1))/baud))
	return
}

// WriteEepromResponseDelay
//
// Writes to the response delay setting stored in EEPROM, changing
// it to the specified value, in units of microseconds.
//
// This function is only useful for Motorons with a UART serial interface,
// and only has an effect if JMP1 is shorted to GND.
//
// **Warning: Be careful not to write to the EEPROM in a fast loop. The
// EEPROM memory of the Motoron's microcontroller is only rated for
// 100,000 erase/write cycles.**
func (se *MXT550) WriteEepromResponseDelay(delay byte) (err error) {
	err = se.WriteEeprom(byte(MotoronSettingResponseDelay), delay)
	return
}

// Reinitialize
//
// Sends a "Reinitialize" command to the Motoron, which resets most of the
// Motoron's variables to their default state.
//
// For more information, see the "Reinitialize" command in the Motoron
// user's guide.
//
// See also: Reset()
func (se *MXT550) Reinitialize() (err error) {

	//
	// Always send the reset command with a CRC byte to make it more reliable.
	//

	err = se.sendCommandCrc([]byte{byte(MotoronCmdReinitialize)})
	if err != nil {
		return
	}

	se.protocolOptions = DefaultProtocolOptions

	time.Sleep(time.Millisecond * 10)
	return
}

// Reset
//
// Sends a "Reset" command to the Motoron, which does a full hardware reset.
//
// This command is equivalent to briefly driving the Motoron's RST pin low.
// The Motoron's RST is briefly driven low by the Motoron itself as a
// result of this command.
//
// After running this command, we recommend waiting for at least 5
// milliseconds before you try to communicate with the Motoron.
//
// See also: Reinitialize()
func (se *MXT550) Reset() (err error) {

	//
	// Always send the reset command with a CRC byte to make it more reliable.
	//

	err = se.sendCommandCrc([]byte{byte(MotoronCmdReset)})
	if err != nil {
		return
	}

	se.protocolOptions = DefaultProtocolOptions

	time.Sleep(time.Millisecond * 10)
	return
}

// GetVariables
//
// Reads information from the Motoron using a "Get variables" command.
//
// This library has helper methods to read every variable, but this method
// is useful if you want to get the raw bytes, or if you want to read
// multiple consecutive variables at the same time for efficiency.
//
// param motor 0 to read general variables, or a motor number to read
// motor-specific variables.
// param offset The location of the first byte to read.
// param length How many bytes to read.
// param buffer A pointer to an array to store the bytes read
// from the controller.
func (se *MXT550) GetVariables(motor byte, offset MotoronVar, length byte) (buffer []byte, err error) {

	cmd := []byte{
		byte(MotoronCmdGetVariables),
		motor & 0x7F,
		byte(offset) & 0x7F,
		length & 0x7F,
	}

	if se.debug {
		if motor == 0 {
			log.Printf("GetVariables name=%s length=%d", MotoronVar2Str[offset], length)
		} else {
			log.Printf("GetVariables motor=%d name=%s length=%d", motor, MotoronMvar2Str[offset], length)
		}
	}

	buffer, err = se.sendCommandAndReadResponse(cmd, length)
	return
}

// GetVar8
//
// Reads one byte from the Motoron using a "Get variables" command.
//
// param motor 0 to read a general variable, or a motor number to read a motor-specific variable.
// param offset The location of the byte to read.
func (se *MXT550) GetVar8(motor byte, offset MotoronVar) (val uint8, err error) {

	buffer, err := se.GetVariables(motor, offset, 1)
	if err != nil {
		return
	}

	val = buffer[0]
	return
}

// GetVar16
//
// Reads two bytes from the Motoron using a "Get variables" command.
//
// param motor 0 to read general variables, or a motor number to read motor-specific variables.
// param offset The location of the first byte to read.
func (se *MXT550) GetVar16(motor byte, offset MotoronVar) (val uint16, err error) {

	buffer, err := se.GetVariables(motor, offset, 2)
	if err != nil {
		return
	}

	val = uint16(buffer[0]) | uint16(buffer[1])<<8
	return
}

// GetProtocolOptions
//
// Reads the "Protocol options" variable from the Motoron.
func (se *MXT550) GetProtocolOptions() (options MotoronProtocolOptions, err error) {

	val, err := se.GetVar8(0, MotoronVarProtocolOptions)
	if err != nil {
		return
	}

	options = MotoronProtocolOptions(val)
	return
}

// GetStatusFlags
//
// Reads the "Status flags" variable from the Motoron.
//
// The bits in this variable are defined by the MOTORON_STATUS_FLAG_*
// macros:
//
// - MOTORON_STATUS_FLAG_PROTOCOL_ERROR
// - MOTORON_STATUS_FLAG_CRC_ERROR
// - MOTORON_STATUS_FLAG_COMMAND_TIMEOUT_LATCHED
// - MOTORON_STATUS_FLAG_MOTOR_FAULT_LATCHED
// - MOTORON_STATUS_FLAG_NO_POWER_LATCHED
// - MOTORON_STATUS_FLAG_RESET
// - MOTORON_STATUS_FLAG_COMMAND_TIMEOUT
// - MOTORON_STATUS_FLAG_MOTOR_FAULTING
// - MOTORON_STATUS_FLAG_NO_POWER
// - MOTORON_STATUS_FLAG_ERROR_ACTIVE
// - MOTORON_STATUS_FLAG_MOTOR_OUTPUT_ENABLED
// - MOTORON_STATUS_FLAG_MOTOR_DRIVING
//
// Here is some example code that uses C++ bitwise operators to check
// whether there is currently a motor fault or a lack of power:
//
// ```{.cpp}
// uint16_t mask = (1 << MOTORON_STATUS_FLAG_NO_POWER) |
//
//	(1 << MOTORON_STATUS_FLAG_MOTOR_FAULTING);
//
// if (getStatusFlags() & mask) {  /* do something */ }
// ```
//
// This library has helper methods that make it easier if you just want to
// read a single bit:
//
// - getProtocolErrorFlag()
// - getCrcErrorFlag()
// - getCommandTimeoutLatchedFlag()
// - getMotorFaultLatchedFlag()
// - getNoPowerLatchedFlag()
// - getResetFlag()
// - getMotorFaultingFlag()
// - getNoPowerFlag()
// - getErrorActiveFlag()
// - getMotorOutputEnabledFlag()
// - getMotorDrivingFlag()
//
// The clearLatchedStatusFlags() method sets the specified set of latched
// status flags to 0.  The reinitialize() and reset() commands reset the
// latched status flags to their default values.
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetStatusFlags() (flags MotoronStatusFlags, err error) {

	val, err := se.GetVar16(0, MotoronVarStatusFlags)
	if err != nil {
		return
	}

	flags = MotoronStatusFlags(val)
	return
}

func (se *MXT550) GetStatusFlagsStrings(flags MotoronStatusFlags) (flagStrings []string) {

	for flag := MotoronStatusFlag(0); flag < 16; flag++ {

		flagStr := MotoronStatusFlag2Str[flag]
		if flagStr == "" {
			continue
		}

		if flags&(1<<flag) == 0 {
			//flagStrings = append(flagStrings, "-"+flagStr)
		} else {
			flagStrings = append(flagStrings, "+"+flagStr)
		}
	}

	return
}

// GetProtocolErrorFlag
//
// Returns the "Protocol error" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetProtocolErrorFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagProtocolError) != 0
	return
}

// GetCrcErrorFlag
//
// Returns the "CRC error" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetCrcErrorFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagCrcError) != 0
	return
}

// GetCommandTimeoutLatchedFlag
//
// Returns the "Command timeout latched" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetCommandTimeoutLatchedFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagCommandTimeoutLatched) != 0
	return
}

// GetMotorFaultLatchedFlag
//
// Returns the "Motor fault latched" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetMotorFaultLatchedFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagMotorFaultLatched) != 0
	return
}

// GetNoPowerLatchedFlag
//
// Returns the "No power latched" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetNoPowerLatchedFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagNoPowerLatched) != 0
	return
}

// GetUARTErrorFlag
//
// Returns the "UART error" bit from getStatusFlags().
//
// This bit is only relevant for the Motoron controllers with a UART serial
// interface.
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
//
// If this flag is set, you might consider calling getUARTFaults()
// to get details about what error happened.
func (se *MXT550) GetUARTErrorFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagUartError) != 0
	return
}

// GetResetFlag
//
// Returns the "Reset" bit from getStatusFlags().
//
// This bit is set to 1 when the Motoron powers on, its processor is
// reset (e.g. by reset()), or it receives a reinitialize() command.
// It can be cleared using clearResetFlag() or clearLatchedStatusFlags().
//
// By default, the Motoron is configured to treat this bit as an error,
// so you will need to clear it before you can turn on the motors.
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetResetFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagReset) != 0
	return
}

// GetCommandTimeoutFlag
//
// Returns the "Command timeout" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetCommandTimeoutFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagCommandTimeout) != 0
	return
}

// GetMotorFaultingFlag
//
// Returns the "Motor faulting" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetMotorFaultingFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagMotorFaulting) != 0
	return
}

// GetNoPowerFlag
//
// Returns the "No power" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetNoPowerFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagNoPower) != 0
	return
}

// GetErrorActiveFlag
//
// Returns the "Error active" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetErrorActiveFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagErrorActive) != 0
	return
}

// GetMotorOutputEnabledFlag
//
// Returns the "Motor output enabled" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetMotorOutputEnabledFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagMotorOutputEnabled) != 0
	return
}

// GetMotorDrivingFlag
//
// Returns the "Motor driving" bit from getStatusFlags().
//
// For more information, see the "Status flags" variable in the Motoron
// user's guide.
func (se *MXT550) GetMotorDrivingFlag() (flag bool, err error) {

	flags, err := se.GetStatusFlags()
	if err != nil {
		return
	}

	flag = flags&(1<<MotoronStatusFlagMotorDriving) != 0
	return
}

// GetUARTFaults
//
// Returns the "UART faults" variable.
//
// Every time the Motoron sets the "UART error" bit in the status flags
// register (see getUARTErrorFlag()) it also sets one of the bits in this
// variable to indicate the cause of the error.
//
// The bits in this variable are defined by the MOTORON_UART_FAULT_*
// macros:
//
// - MOTORON_UART_FAULT_FRAMING
// - MOTORON_UART_FAULT_NOISE
// - MOTORON_UART_FAULT_HARDWARE_OVERRUN
// - MOTORON_UART_FAULT_SOFTWARE_OVERRUN
//
// This function is only useful for Motorons with a UART serial interface.
//
// For more information, see the "UART faults" variable in the Motoron
// user's guide.
//
// See also: ClearUARTFaults()
func (se *MXT550) GetUARTFaults() (val uint8, err error) {
	val, err = se.GetVar8(0, MotoronVarUartFaults)
	return
}

// GetVinVoltage
//
// Reads voltage on the Motoron's VIN pin, in raw device units.
//
// For more information, see the "VIN voltage" variable in the Motoron
// user's guide.
//
// See also: GetVinVoltageMv()
func (se *MXT550) GetVinVoltage() (val uint16, err error) {
	val, err = se.GetVar16(0, MotoronVarVinVoltage)
	return
}

// GetVinVoltageMv
//
// Reads the voltage on the Motoron's VIN pin and converts it to milli-volts.
//
// For more information, see the "VIN voltage" variable in the Motoron
// user's guide.
//
// param referenceMv The logic voltage of the Motoron, in milli-volts.
// param mType Specifies what type of Motoron you are using.
//
// See also: GetVinVoltage()
func (se *MXT550) GetVinVoltageMv(referenceMv uint16, mType MotoronVinSenseType) (val uint32, err error) {

	scale := uint32(1047)
	if mType&1 == 1 {
		scale = 459
	}

	tmpVal, err := se.GetVinVoltage()
	if err != nil {
		return
	}

	val = uint32(tmpVal) * uint32(referenceMv) / 1024 * scale / 47
	return
}

// GetCommandTimeoutMilliseconds
//
// Reads the "Command timeout" variable and converts it to milliseconds.
//
// For more information, see the "Command timeout" variable in the Motoron
// user's guide.
//
// See also: SetCommandTimeoutMilliseconds()
func (se *MXT550) GetCommandTimeoutMilliseconds() (val uint16, err error) {
	val, err = se.GetVar16(0, MotoronVarCommandTimeout)
	val *= 4
	return
}

// GetErrorResponse
//
// Reads the "Error response" variable, which defines how the Motoron will
// stop its motors when an error is happening.
//
// For more information, see the "Error response" variable in the Motoron
// user's guide.
//
// See also: SetErrorResponse()
func (se *MXT550) GetErrorResponse() (val uint8, err error) {
	val, err = se.GetVar8(0, MotoronVarErrorResponse)
	return
}

// GetErrorMask
//
// Reads the "Error mask" variable, which defines which status flags are
// considered to be errors.
//
// For more information, see the "Error mask" variable in the Motoron
// user's guide.
//
// See also: SetErrorMask()
func (se *MXT550) GetErrorMask() (val uint16, err error) {
	val, err = se.GetVar16(0, MotoronVarErrorMask)
	return
}

// GetJumperState
//
// Reads the "Jumper state" variable.
//
// For more information, see the "Jumper state" variable in the Motoron
// user's guide.
func (se *MXT550) GetJumperState() (val uint8, err error) {
	val, err = se.GetVar8(0, MotoronVarJumperState)
	return
}

// GetTargetSpeed
//
// Reads the target speed of the specified motor, which is the speed at
// which the motor has been commanded to move.
//
// For more information, see the "Target speed" variable in the Motoron
// user's guide.
//
// See also: SetSpeed(), SetAllSpeeds(), SetAllSpeedsUsingBuffers()
func (se *MXT550) GetTargetSpeed(motor byte) (speed int16, err error) {
	val, err := se.GetVar16(motor, MotoronMvarTargetSpeed)
	speed = int16(val)
	return
}

// GetTargetBrakeAmount
//
// Reads the target brake amount for the specified motor.
//
// For more information, see the "Target speed" variable in the Motoron
// user's guide.
//
// See also: SetTargetBrakeAmount()
func (se *MXT550) GetTargetBrakeAmount(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarTargetBrakeAmount)
	return
}

// GetCurrentSpeed
//
// Reads the current speed of the specified motor, which is the speed that
// the Motoron is currently trying to apply to the motor.
//
// For more information, see the "Target speed" variable in the Motoron
// user's guide.
//
// See also: SetSpeedNow(), SetAllSpeedsNow(), SetAllSpeedsNowUsingBuffers()
func (se *MXT550) GetCurrentSpeed(motor byte) (speed int16, err error) {
	val, err := se.GetVar16(motor, MotoronMvarCurrentSpeed)
	speed = int16(val)
	return
}

// GetBufferedSpeed
//
// Reads the buffered speed of the specified motor.
//
// For more information, see the "Buffered speed" variable in the Motoron
// user's guide.
//
// See also: SetBufferedSpeed(), SetAllBufferedSpeeds()
func (se *MXT550) GetBufferedSpeed(motor byte) (speed int16, err error) {
	val, err := se.GetVar16(motor, MotoronMvarBufferedSpeed)
	speed = int16(val)
	return
}

// GetPwmMode
//
// Reads the PWM mode of the specified motor.
//
// For more information, see the "PWM mode" variable in the Motoron
// user's guide.
//
// See also: SetPwmMode()
func (se *MXT550) GetPwmMode(motor byte) (val uint8, err error) {
	val, err = se.GetVar8(motor, MotoronMvarPwmMode)
	return
}

// GetMaxAccelerationForward
//
// Reads the maximum acceleration of the specified motor for the forward
// direction.
//
// For more information, see the "Max acceleration forward" variable in the
// Motoron user's guide.
//
// See also: SetMaxAcceleration(), SetMaxAccelerationForward()
func (se *MXT550) GetMaxAccelerationForward(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarMaxAccelForward)
	return
}

// GetMaxAccelerationReverse
//
// Reads the maximum acceleration of the specified motor for the reverse
// direction.
//
// For more information, see the "Max acceleration reverse" variable in the
// Motoron user's guide.
//
// See also: SetMaxAcceleration(), SetMaxAccelerationReverse()
func (se *MXT550) GetMaxAccelerationReverse(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarMaxAccelReverse)
	return
}

// GetMaxDecelerationForward
//
// Reads the maximum deceleration of the specified motor for the forward
// direction.
//
// For more information, see the "Max deceleration forward" variable in the
// Motoron user's guide.
//
// See also: SetMaxDeceleration(), SetMaxDecelerationForward()
func (se *MXT550) GetMaxDecelerationForward(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarMaxDecelForward)
	return
}

// GetMaxDecelerationReverse
//
// Reads the maximum deceleration of the specified motor for the reverse
// direction.
//
// For more information, see the "Max deceleration reverse" variable in the
// Motoron user's guide.
//
// See also: SetMaxDeceleration(), SetMaxDecelerationReverse()
func (se *MXT550) GetMaxDecelerationReverse(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarMaxDecelReverse)
	return
}

// GetMaxDecelerationTemporary
//
// This function is used by Pololu for testing.
func (se *MXT550) GetMaxDecelerationTemporary(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarMaxDecelTemporary)
	return
}

// GetStartingSpeedForward
//
// Reads the starting speed for the specified motor in the forward direction.
//
// For more information, see the "Starting speed forward" variable in the
// Motoron user's guide.
//
// See also: SetStartingSpeed(), SetStartingSpeedForward()
func (se *MXT550) GetStartingSpeedForward(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarStartingSpeedForward)
	return
}

// GetStartingSpeedReverse
//
// Reads the starting speed for the specified motor in the reverse direction.
//
// For more information, see the "Starting speed reverse" variable in the
// Motoron user's guide.
//
// See also: SetStartingSpeed(), SetStartingSpeedReverse()
func (se *MXT550) GetStartingSpeedReverse(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarStartingSpeedReverse)
	return
}

// GetDirectionChangeDelayForward
//
// Reads the direction change delay for the specified motor in the
// forward direction.
//
// For more information, see the "Direction change delay forward" variable
// in the Motoron user's guide.
//
// See also: SetDirectionChangeDelay(), SetDirectionChangeDelayForward()
func (se *MXT550) GetDirectionChangeDelayForward(motor byte) (val uint8, err error) {
	val, err = se.GetVar8(motor, MotoronMvarDirectionChangeDelayForward)
	return
}

// GetDirectionChangeDelayReverse
//
// Reads the direction change delay for the specified motor in the
// reverse direction.
//
// For more information, see the "Direction change delay reverse" variable
// in the Motoron user's guide.
//
// See also: SetDirectionChangeDelay(), SetDirectionChangeDelayReverse()
func (se *MXT550) GetDirectionChangeDelayReverse(motor byte) (val uint8, err error) {
	val, err = se.GetVar8(motor, MotoronMvarDirectionChangeDelayReverse)
	return
}

// GetCurrentLimit
//
// Reads the current limit for the specified motor.
//
// This only works for the high-power Motorons.
//
// For more information, see the "Current limit" variable
// in the Motoron user's guide.
//
// See also: SetCurrentLimit()
func (se *MXT550) GetCurrentLimit(motor byte) (val uint16, err error) {
	val, err = se.GetVar16(motor, MotoronMvarCurrentLimit)
	return
}

// GetCurrentSenseReading
//
// Reads all the results from the last current sense measurement for the
// specified motor.
//
// This function reads the "Current sense raw", "Current sense speed", and
// "Current sense processed" variables from the Motoron using a single
// command, so the values returned are all guaranteed to be part of the
// same measurement.
//
// This only works for the high-power Motorons.
//
// See also: GetCurrentSenseRawAndSpeed(), GetCurrentSenseProcessedAndSpeed()
func (se *MXT550) GetCurrentSenseReading(motor byte) (mcsr *MotoronCurrentSenseReading, err error) {

	buffer, err := se.GetVariables(motor, MotoronMvarCurrentSenseRaw, 6)
	if err != nil {
		return
	}

	mcsr = &MotoronCurrentSenseReading{
		Raw:       uint16(buffer[0]) | uint16(buffer[1])<<8,
		Speed:     int16(buffer[2]) | int16(buffer[3])<<8,
		Processed: uint16(buffer[4]) | uint16(buffer[5])<<8,
	}

	return
}

// GetCurrentSenseRawAndSpeed
//
// This is like getCurrentSenseReading() but it only reads the raw current
// sense measurement and the speed.
//
// The 'processed' member of the returned structure will be 0.
//
// This only works for the high-power Motorons.
func (se *MXT550) GetCurrentSenseRawAndSpeed(motor byte) (mcsr *MotoronCurrentSenseReading, err error) {

	buffer, err := se.GetVariables(motor, MotoronMvarCurrentSenseRaw, 4)
	if err != nil {
		return
	}

	mcsr = &MotoronCurrentSenseReading{
		Raw:       uint16(buffer[0]) | uint16(buffer[1])<<8,
		Speed:     int16(buffer[2]) | int16(buffer[3])<<8,
		Processed: uint16(buffer[4]) | uint16(buffer[5])<<8,
	}

	return
}

// GetCurrentSenseProcessedAndSpeed
//
// This is like getCurrentSenseReading() but it only reads the processed
// current sense measurement and the speed.
//
// The 'raw' member of the returned structure will be 0.
//
// This only works for the high-power Motorons.
func (se *MXT550) GetCurrentSenseProcessedAndSpeed(motor byte) (mcsr *MotoronCurrentSenseReading, err error) {

	buffer, err := se.GetVariables(motor, MotoronMvarCurrentSenseSpeed, 4)
	if err != nil {
		return
	}

	mcsr = &MotoronCurrentSenseReading{
		Raw:       uint16(buffer[0]) | uint16(buffer[1])<<8,
		Processed: uint16(buffer[2]) | uint16(buffer[3])<<8,
	}

	return
}

// GetCurrentSenseRaw
//
// Reads the raw current sense measurement for the specified motor.
//
// This only works for the high-power Motorons.
//
// For more information, see the "Current sense raw" variable
// in the Motoron user's guide.
//
// See also: GetCurrentSenseReading()
func (se *MXT550) GetCurrentSenseRaw(motor byte) (senseRaw uint16, err error) {
	senseRaw, err = se.GetVar16(motor, MotoronMvarCurrentSenseRaw)
	return
}

// GetCurrentSenseProcessed
//
// Reads the processed current sense reading for the specified motor.
//
// This only works for the high-power Motorons.
//
// The units of this reading depend on the logic voltage of the Motoron
// and on the specific model of Motoron that you have, and you can use
// MotoronI2C::currentSenseUnitsMilliamps() to calculate the units.
//
// The accuracy of this reading can be improved by measuring the current
// sense offset and setting it with setCurrentSenseOffset().
// See the "Current sense processed" variable in the Motoron user's guide for
// or the CurrentSenseCalibrate example for more information.
//
// Note that this reading will be 0xFFFF if an overflow happens during the
// calculation due to very high current.
//
// See also: GetCurrentSenseProcessedAndSpeed()
func (se *MXT550) GetCurrentSenseProcessed(motor byte) (senseProcessed uint16, err error) {
	senseProcessed, err = se.GetVar16(motor, MotoronMvarCurrentSenseProcessed)
	return
}

// GetCurrentSenseOffset
//
// Reads the current sense offset setting.
//
// This only works for the high-power Motorons.
//
// For more information, see the "Current sense offset" variable in the
// Motoron user's guide.
//
// See also: SetCurrentSenseOffset()
func (se *MXT550) GetCurrentSenseOffset(motor byte) (senseOffset uint8, err error) {
	senseOffset, err = se.GetVar8(motor, MotoronMvarCurrentSenseOffset)
	return
}

// GetCurrentSenseMinimumDivisor
//
// Reads the current sense minimum divisor setting and returns it as a speed
// between 0 and 800.
//
// This only works for the high-power Motorons.
//
// For more information, see the "Current sense minimum divisor" variable in
// the Motoron user's guide.
//
// See also: SetCurrentSenseMinimumDivisor()
func (se *MXT550) GetCurrentSenseMinimumDivisor(motor byte) (minimumDivisor uint16, err error) {
	temp, err := se.GetVar8(motor, MotoronMvarCurrentSenseOffset)
	minimumDivisor = uint16(temp) << 2
	return
}

// SetVariable
//
// Configures the Motoron using a "Set variable" command.
//
// This library has helper methods to set every variable, so you should
// not need to call this function directly.
//
// param motor 0 to set a general variable, or a motor number to set
//
//	motor-specific variables.
//
// param offset The address of the variable to set (only certain offsets
//
//	are allowed).
//
// param value The value to set the variable to.
//
// See also: GetVariables()
func (se *MXT550) SetVariable(motor byte, offset MotoronVar, value uint16) (err error) {

	if value > 0x3fff {
		value = 0x3fff
	}

	cmd := []byte{
		byte(MotoronCmdSetVariable),
		motor & 0x1f,
		byte(offset & 0x7f),
		byte(value & 0x7f),
		byte(value >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// SetCommandTimeoutMilliseconds
//
// Sets the command timeout period, in milliseconds.
//
// For more information, see the "Command timeout" variable
// in the Motoron user's guide.
//
// See also: DisableCommandTimeout(), GetCommandTimeoutMilliseconds()
func (se *MXT550) SetCommandTimeoutMilliseconds(ms uint16) (err error) {

	// Divide by 4, but round up, and make sure we don't have
	// an overflow if 0xFFFF is passed.

	timeout := ms / 4
	if ms&3 != 0 {
		timeout++
	}

	err = se.SetVariable(0, MotoronVarCommandTimeout, timeout)
	return
}

// SetErrorResponse
//
// Sets the error response, which defines how the Motoron will
// stop its motors when an error is happening.
//
// The response parameter should be one of:
//
// - MOTORON_ERROR_RESPONSE_COAST
// - MOTORON_ERROR_RESPONSE_BRAKE
// - MOTORON_ERROR_RESPONSE_COAST_NOW
// - MOTORON_ERROR_RESPONSE_BRAKE_NOW
//
// For more information, see the "Error response" variable in the Motoron
// user's guide.
//
// See also: GetErrorResponse()
func (se *MXT550) SetErrorResponse(response uint8) (err error) {
	err = se.SetVariable(0, MotoronVarErrorResponse, uint16(response))
	return
}

// SetErrorMask
//
// Sets the "Error mask" variable, which defines which status flags are
// considered to be errors.
//
// For more information, see the "Error mask" variable in the Motoron
// user's guide.
//
// See also: GetErrorMask(), GetStatusFlags()
func (se *MXT550) SetErrorMask(mask MotoronStatusFlags) (err error) {
	err = se.SetVariable(0, MotoronVarErrorMask, uint16(mask))
	return
}

// DisableCommandTimeout
//
// This disables the Motoron's command timeout feature by resetting
// the "Error mask" variable to its default value but with the command
// timeout bit cleared.
//
// By default, the Motoron's command timeout will occur if no valid commands
// are received in 1500 milliseconds, and the command timeout is treated as
// an error, so the motors will shut down.  You can use this function if you
// want to disable that feature.
//
// Note that this function overrides any previous values you set in the
// "Error mask" variable, so if you are using SetErrorMask() in your program
// to configure which status flags are treated as errors, you do not need to
// use this function, and you probably should not use this function.
//
// See also: SetCommandTimeoutMilliseconds(), SetErrorMask()
func (se *MXT550) DisableCommandTimeout() (err error) {
	err = se.SetErrorMask(DefaultErrorMask ^ (1 << MotoronStatusFlagCommandTimeout))
	return
}

// ClearUARTFaults
//
// Sends a "Set variable" command that clears the specified flags in
// GetUARTFaults().
//
// For each bit in the flags argument that is 1, this command clears the
// corresponding bit in the "UART faults" variable, setting it to 0.
//
// For more information, see the "UART faults" variable in the
// Motoron user's guide.
func (se *MXT550) ClearUARTFaults(flags uint8) (err error) {
	err = se.SetVariable(0, MotoronVarUartFaults, 0x3FFF^uint16(flags))
	return
}

// SetPwmMode
//
// Sets the PWM mode for the specified motor.
//
// The mode parameter should be one of the following:
//
// - MOTORON_PWM_MODE_DEFAULT (20 kHz)
// - MOTORON_PWM_MODE_1_KHZ 1
// - MOTORON_PWM_MODE_2_KHZ 2
// - MOTORON_PWM_MODE_4_KHZ 3
// - MOTORON_PWM_MODE_5_KHZ 4
// - MOTORON_PWM_MODE_10_KHZ 5
// - MOTORON_PWM_MODE_20_KHZ 6
// - MOTORON_PWM_MODE_40_KHZ 7
// - MOTORON_PWM_MODE_80_KHZ 8
//
// For more information, see the "PWM mode" variable in the Motoron user's
// guide.
//
// See also: GetPwmMode()
func (se *MXT550) SetPwmMode(motor byte, mode uint8) (err error) {
	err = se.SetVariable(motor, MotoronMvarPwmMode, uint16(mode))
	return
}

// SetMaxAccelerationForward
//
// Sets the maximum acceleration of the specified motor for the forward
// direction.
//
// For more information, see the "Max acceleration forward" variable in the
// Motoron user's guide.
//
// See also: SetMaxAcceleration(), GetMaxAccelerationForward()
func (se *MXT550) SetMaxAccelerationForward(motor byte, accel uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarMaxAccelForward, accel)
	return
}

// SetMaxAccelerationReverse
//
// Sets the maximum acceleration of the specified motor for the reverse
// direction.
//
// For more information, see the "Max acceleration reverse" variable in the
// Motoron user's guide.
//
// See also: SetMaxAcceleration(), GetMaxAccelerationReverse()
func (se *MXT550) SetMaxAccelerationReverse(motor byte, accel uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarMaxAccelReverse, accel)
	return
}

// SetMaxAcceleration
//
// Sets the maximum acceleration of the specified motor (both directions).
//
// If this function succeeds, it is equivalent to calling
// SetMaxAccelerationForward() and SetMaxAccelerationReverse().
func (se *MXT550) SetMaxAcceleration(motor byte, accel uint16) (err error) {

	err = se.SetMaxAccelerationForward(motor, accel)
	if err != nil {
		return
	}

	err = se.SetMaxAccelerationReverse(motor, accel)
	return
}

// SetMaxDecelerationForward
//
// Sets the maximum deceleration of the specified motor for the forward
// direction.
//
// For more information, see the "Max deceleration forward" variable in the
// Motoron user's guide.
//
// See also: SetMaxDeceleration(), GetMaxDecelerationForward()
func (se *MXT550) SetMaxDecelerationForward(motor byte, decel uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarMaxDecelForward, decel)
	return
}

// SetMaxDecelerationReverse
//
// Sets the maximum deceleration of the specified motor for the reverse
// direction.
//
// For more information, see the "Max deceleration reverse" variable in the
// Motoron user's guide.
//
// See also: SetMaxDeceleration(), GetMaxDecelerationReverse()
func (se *MXT550) SetMaxDecelerationReverse(motor byte, decel uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarMaxDecelReverse, decel)
	return
}

// SetMaxDeceleration
//
// Sets the maximum deceleration of the specified motor (both directions).
//
// If this function succeeds, it is equivalent to calling
// SetMaxDecelerationForward() and SetMaxDecelerationReverse().
func (se *MXT550) SetMaxDeceleration(motor byte, decel uint16) (err error) {

	err = se.SetMaxDecelerationForward(motor, decel)
	if err != nil {
		return
	}

	err = se.SetMaxDecelerationReverse(motor, decel)
	return
}

// SetStartingSpeedForward
//
// Sets the starting speed of the specified motor for the forward
// direction.
//
// For more information, see the "Starting speed forward" variable in the
// Motoron user's guide.
//
// See also: SetStartingSpeed(), GetStartingSpeedForward()
func (se *MXT550) SetStartingSpeedForward(motor byte, speed uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarStartingSpeedForward, speed)
	return
}

// SetStartingSpeedReverse
//
// Sets the starting speed of the specified motor for the reverse
// direction.
//
// For more information, see the "Starting speed reverse" variable in the
// Motoron user's guide.
//
// See also: SetStartingSpeed(), GetStartingSpeedReverse()
func (se *MXT550) SetStartingSpeedReverse(motor byte, speed uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarStartingSpeedReverse, speed)
	return
}

// SetStartingSpeed
//
// Sets the starting speed of the specified motor (both directions).
//
// If this function succeeds, it is equivalent to calling
// SetStartingSpeedForward() and SetStartingSpeedReverse().
func (se *MXT550) SetStartingSpeed(motor byte, speed uint16) (err error) {

	err = se.SetStartingSpeedForward(motor, speed)
	if err != nil {
		return
	}

	err = se.SetStartingSpeedReverse(motor, speed)
	return
}

// SetDirectionChangeDelayForward
//
// Sets the direction change delay of the specified motor for the forward
// direction, in units of 10 ms.
//
// For more information, see the "Direction change delay forward" variable
// in the Motoron user's guide.
//
// See also: SetDirectionChangeDelay(), GetDirectionChangeDelayForward()
func (se *MXT550) SetDirectionChangeDelayForward(motor byte, duration uint8) (err error) {
	err = se.SetVariable(motor, MotoronMvarDirectionChangeDelayForward, uint16(duration))
	return
}

// SetDirectionChangeDelayReverse
//
// Sets the direction change delay of the specified motor for the reverse
// direction, in units of 10 ms.
//
// For more information, see the "Direction change delay reverse" variable
// in the Motoron user's guide.
//
// See also: SetDirectionChangeDelay(), GetDirectionChangeDelayReverse()
func (se *MXT550) SetDirectionChangeDelayReverse(motor byte, duration uint8) (err error) {
	err = se.SetVariable(motor, MotoronMvarDirectionChangeDelayReverse, uint16(duration))
	return
}

// SetDirectionChangeDelay
//
// Sets the direction change delay of the specified motor (both directions),
// in units of 10 ms.
//
// If this function succeeds, it is equivalent to calling
// SetDirectionChangeDelayForward() and SetDirectionChangeDelayReverse().
func (se *MXT550) SetDirectionChangeDelay(motor byte, duration uint8) (err error) {

	err = se.SetDirectionChangeDelayForward(motor, duration)
	if err != nil {
		return
	}

	err = se.SetDirectionChangeDelayReverse(motor, duration)
	return
}

// SetCurrentLimit
//
// Sets the current limit for the specified motor.
//
// This only works for the high-power Motorons.
//
// The units of the current limit depend on the type of Motoron you have
// and the logic voltage of your system.  See the "Current limit" variable
// in the Motoron user's guide for more information, or see
// MotoronI2C::calculateCurrentLimit().
//
// See also: GetCurrentLimit()
func (se *MXT550) SetCurrentLimit(motor byte, limit uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarCurrentLimit, limit)
	return
}

// SetCurrentSenseOffset
//
// Sets the current sense offset setting for the specified motor.
//
// This is one of the settings that determines how current sense
// readings are processed.  It is supposed to be the value returned by
// GetCurrentSenseRaw() when Motor power is supplied to the Motoron, and
// it is driving its motor outputs at speed 0.
//
// The CurrentSenseCalibrate example shows how to measure the current
// sense offsets and load them onto the Motoron using this function.
//
// If you do not care about measuring motor current, you do not need to
// set this variable.
//
// For more information, see the "Current sense offset" variable in the
// Motoron user's guide.
//
// This only works for the high-power Motorons.
//
// See also: GetCurrentSenseOffset()
func (se *MXT550) SetCurrentSenseOffset(motor byte, offset uint8) (err error) {
	err = se.SetVariable(motor, MotoronMvarCurrentSenseOffset, uint16(offset))
	return
}

// SetCurrentSenseMinimumDivisor
//
// Sets the current sense minimum divisor setting for the specified motor,
// given a speed between 0 and 800.
//
// This is one of the settings that determines how current sense
// readings are processed.
//
// If you do not care about measuring motor current, you do not need to
// set this variable.
//
// For more information, see the "Current sense minimum divisor" variable in
// the Motoron user's guide.
//
// This only works for the high-power Motorons.
//
// See also: GetCurrentSenseMinimumDivisor()
func (se *MXT550) SetCurrentSenseMinimumDivisor(motor byte, speed uint16) (err error) {
	err = se.SetVariable(motor, MotoronMvarCurrentSenseMinimumDivisor, speed>>2)
	return
}

// CoastNow
//
// Sends a "Coast now" command to the Motoron, causing all motors to
// immediately start coasting.
//
// For more information, see the "Coast now" command in the Motoron
// user's guide.
func (se *MXT550) CoastNow() (err error) {

	cmd := []byte{
		byte(MotoronCmdCoastNow),
	}

	err = se.sendCommand(cmd)
	return
}

// ClearMotorFault
//
// Sends a "Clear motor fault" command to the Motoron.
//
// If any of the Motoron's motors chips are currently experiencing a
// fault (error), or bit 0 of the flags argument is 1, this command makes
// the Motoron attempt to recover from the faults.
//
// For more information, see the "Clear motor fault" command in the Motoron
// user's guide.
//
// See also: ClearMotorFaultUnconditional(), GetMotorFaultingFlag()
func (se *MXT550) ClearMotorFault(flags uint8) (err error) {

	cmd := []byte{
		byte(MotoronCmdClearMotorFault),
		flags & 0x7f,
	}

	err = se.sendCommand(cmd)
	return
}

// ClearMotorFaultUnconditional
//
// Sends a "Clear motor fault" command to the Motoron with the
// "unconditional" flag set, so the Motoron will attempt to recover
// from any motor faults even if no fault is currently occurring.
//
// This is a more robust version of clearMotorFault().
func (se *MXT550) ClearMotorFaultUnconditional() (err error) {
	err = se.ClearMotorFault(1 << MotoronClearMotorFaultUnconditional)
	return
}

// ClearLatchedStatusFlags
//
// Clears the specified flags in GetStatusFlags().
//
// For each bit in the flags argument that is 1, this command clears the
// corresponding bit in the "Status flags" variable, setting it to 0.
//
// For more information, see the "Clear latched status flags" command in the
// Motoron user's guide.
//
// See also: GetStatusFlags(), SetLatchedStatusFlags()
func (se *MXT550) ClearLatchedStatusFlags(flags uint16) (err error) {

	cmd := []byte{
		byte(MotoronCmdClearLatchedStatusFlags),
		byte(flags & 0x7f),
		byte(flags >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// ClearResetFlag
//
// Clears the Motoron's reset flag.
//
// The reset flag is a latched status flag in GetStatusFlags() that is
// particularly important to clear: it gets set to 1 after the Motoron
// powers on or experiences a reset, and it is considered to be an error
// by default, so it prevents the motors from running.  Therefore, it is
// necessary to call this function (or clearLatchedStatusFlags()) to clear
// the Reset flag before you can get the motors running.
//
// We recommend that immediately after you clear the reset flag. you should
// configure the Motoron's motor settings and error response settings.
// That way, if the Motoron experiences an unexpected reset while your system
// is running, it will stop running its motors, and it will not start them
// again until all the important settings have been configured.
//
// See also: ClearLatchedStatusFlags()
func (se *MXT550) ClearResetFlag() (err error) {
	err = se.ClearLatchedStatusFlags(1 << MotoronStatusFlagReset)
	return
}

// SetLatchedStatusFlags
//
// Sets the specified flags in GetStatusFlags().
//
// For each bit in the flags argument that is 1, this command sets the
// corresponding bit in the "Status flags" variable to 1.
//
// For more information, see the "Set latched status flags" command in the
// Motoron user's guide.
//
// See also: GetStatusFlags(), SetLatchedStatusFlags()
func (se *MXT550) SetLatchedStatusFlags(flags uint16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetLatchedStatusFlags),
		byte(flags & 0x7f),
		byte(flags >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// SetSpeed
//
// Sets the target speed of the specified motor.
//
// The current speed will start moving to the specified target speed,
// obeying any acceleration and deceleration limits.
//
// The motor number should be between 1 and the number of motors supported
// by the Motoron.
//
// The speed should be between -800 and 800.  Values outside that range
// will be clipped to -800 or 800 by the Motoron firmware.
//
// For single-channel Motorons, it is better to use SetAllSpeeds() instead
// of this, since it sends one fewer byte.
//
// For more information, see the "Set speed" command in the Motoron
// user's guide.
//
// See also: SetSpeedNow(), SetAllSpeeds()
func (se *MXT550) SetSpeed(motor byte, speed int16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetSpeed),
		motor & 0x7f,
		byte(speed & 0x7f),
		byte(speed >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// SetSpeedNow
//
// Sets the target and current speed of the specified motor, ignoring
// any acceleration and deceleration limits.
//
// For single-channel Motorons, it is better to use SetAllSpeedsNow() instead
// of this, since it sends one fewer byte.
//
// For more information, see the "Set speed" command in the Motoron
// user's guide.
//
// See also: SetSpeed(), SetAllSpeedsNow()
func (se *MXT550) SetSpeedNow(motor byte, speed int16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetSpeedNow),
		motor & 0x7f,
		byte(speed & 0x7f),
		byte(speed >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// SetBufferedSpeed
//
// Sets the buffered speed of the specified motor.
//
// This command does not immediately cause any change to the motor: it
// stores a speed for the specified motor in the Motoron so it can be
// used by later commands.
//
// For single-channel Motorons, it is better to use SetAllBufferedSpeeds()
// instead of this, since it sends one fewer byte.
//
// For more information, see the "Set speed" command in the Motoron
// user's guide.
//
// See also: SetSpeed(), SetAllBufferedSpeeds(),
// SetAllSpeedsUsingBuffers(), SetAllSpeedsNowUsingBuffers()
func (se *MXT550) SetBufferedSpeed(motor byte, speed int16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetBufferedSpeed),
		motor & 0x7f,
		byte(speed & 0x7f),
		byte(speed >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// SetAllSpeeds
//
// Sets the target speeds of all the motors at the same time.
//
// The number of speed arguments you provide to this function must be equal
// to the number of motor channels your Motoron has, or else this command
// might not work.
//
// This is equivalent to calling SetSpeed() once for each motor, but it is
// more efficient because all speeds are sent in the same command.
//
// For more information, see the "Set all speeds" command in the Motoron
// user's guide.
//
// See also: SetSpeed(), SetAllSpeedsNow(), SetAllBufferedSpeeds()
func (se *MXT550) SetAllSpeeds(speeds ...int16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetAllSpeeds),
	}

	for _, speed := range speeds {
		cmd = append(cmd, byte(speed&0x7f), byte(speed>>7&0x7f))
	}

	err = se.sendCommand(cmd)
	return
}

// SetAllSpeedsNow
//
// Sets the target and current speeds of all the motors at the same time.
//
// The number of speed arguments you provide to this function must be equal
// to the number of motor channels your Motoron has, or else this command
// might not work.
//
// This is equivalent to calling SetSpeedNow() once for each motor, but it is
// more efficient because all speeds are sent in the same command.
//
// For more information, see the "Set all speeds" command in the Motoron
// user's guide.
//
// See also: SetSpeed(), SetSpeedNow(), SetAllSpeeds()
func (se *MXT550) SetAllSpeedsNow(speeds ...int16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetAllSpeedsNow),
	}

	for _, speed := range speeds {
		cmd = append(cmd, byte(speed&0x7f), byte(speed>>7&0x7f))
	}

	err = se.sendCommand(cmd)
	return
}

// SetAllBufferedSpeeds
//
// Sets the buffered speeds of all the motors.
//
// The number of speed arguments you provide to this function must be equal
// to the number of motor channels your Motoron has, or else this command
// might not work.
//
// This command does not immediately cause any change to the motors: it
// stores speed for each motor in the Motoron so they can be used by later
// commands.
//
// This is equivalent to calling SetBufferedSpeed() once for each motor,
// but it is more efficient because all speeds are sent in the same
// command.
//
// For more information, see the "Set all speeds" command in the Motoron
// user's guide.
//
// See also: SetSpeed(), SetBufferedSpeed(), SetAllSpeeds(),
// SetAllSpeedsUsingBuffers(), SetAllSpeedsNowUsingBuffers()
func (se *MXT550) SetAllBufferedSpeeds(speeds ...int16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetAllBufferedSpeeds),
	}

	for _, speed := range speeds {
		cmd = append(cmd, byte(speed&0x7f), byte(speed>>7&0x7f))
	}

	err = se.sendCommand(cmd)
	return
}

// SetAllSpeedsUsingBuffers
//
// Sets each motor's target speed equal to the buffered speed.
//
// This command is the same as SetAllSpeeds() except that the speeds are
// provided ahead of time using SetBufferedSpeed() or SetAllBufferedSpeeds().
//
// See also: SetAllSpeedsNowUsingBuffers(), SetBufferedSpeed(),
//
//	setAllBufferedSpeeds()
func (se *MXT550) SetAllSpeedsUsingBuffers() (err error) {

	cmd := []byte{
		byte(MotoronCmdSetAllSpeedsUsingBuffers),
	}

	err = se.sendCommand(cmd)
	return
}

// SetAllSpeedsNowUsingBuffers
//
// Sets each motor's target speed and current speed equal to the buffered
// speed.
//
// This command is the same as SetAllSpeedsNow() except that the speeds are
// provided ahead of time using SetBufferedSpeed() or SetAllBufferedSpeeds().
//
// See also: SetAllSpeedsUsingBuffers(), SetBufferedSpeed(), SetAllBufferedSpeeds()
func (se *MXT550) SetAllSpeedsNowUsingBuffers() (err error) {

	cmd := []byte{
		byte(MotoronCmdSetAllSpeedsNowUsingBuffers),
	}

	err = se.sendCommand(cmd)
	return
}

// SetBraking
//
// Commands the motor to brake, coast, or something in between.
//
// Sending this command causes the motor to decelerate to speed 0 obeying
// any relevant deceleration limits.  Once the current speed reaches 0, the
// motor will attempt to brake or coast as specified by this command, but
// due to hardware limitations it might not be able to.
//
// The motor number parameter should be between 1 and the number of motors
// supported by the Motoron.
//
// The amount parameter gets stored in the "Target brake amount" variable
// for the motor and should be between 0 (coasting) and 800 (braking).
// Values above 800 will be clipped to 800 by the Motoron firmware.
//
// See the "Set braking" command in the Motoron user's guide for more
// information.
//
// See also: SetBrakingNow(), GetTargetBrakeAmount()
func (se *MXT550) SetBraking(motor byte, amount uint16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetBraking),
		motor & 0x7f,
		byte(amount & 0x7f),
		byte(amount >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// SetBrakingNow
//
// Commands the motor to brake, coast, or something in between.
//
// Sending this command causes the motor's current speed to change to 0.
// The motor will attempt to brake or coast as specified by this command,
// but due to hardware limitations it might not be able to.
//
// The motor number parameter should be between 1 and the number of motors
// supported by the Motoron.
//
// The amount parameter gets stored in the "Target brake amount" variable
// for the motor and should be between 0 (coasting) and 800 (braking).
// Values above 800 will be clipped to 800 by the Motoron firmware.
//
// See the "Set braking" command in the Motoron user's guide for more
// information.
//
// See also: SetBraking(), GetTargetBrakeAmount()
func (se *MXT550) SetBrakingNow(motor byte, amount uint16) (err error) {

	cmd := []byte{
		byte(MotoronCmdSetBrakingNow),
		motor & 0x7f,
		byte(amount & 0x7f),
		byte(amount >> 7 & 0x7f),
	}

	err = se.sendCommand(cmd)
	return
}

// ResetCommandTimeout
//
// Resets the command timeout.
//
// This prevents the command timeout status flags from getting set for some
// time.  (The command timeout is also reset by every other Motoron command,
// as long as its parameters are valid.)
//
// For more information, see the "Reset command timeout" command in the
// Motoron user's guide.
//
// See also: DisableCommandTimeout(), SetCommandTimeoutMilliseconds()
func (se *MXT550) ResetCommandTimeout() (err error) {

	cmd := []byte{
		byte(MotoronCmdResetCommandTimeout),
	}

	err = se.sendCommand(cmd)
	return
}

// CalculateCurrentLimit
//
// Calculates a current limit value that can be passed to the Motoron
// using setCurrentLimit().
//
// param milliamps The desired current limit, in units of mA.
// param type Specifies what type of Motoron you are using.
// param referenceMv The reference voltage (IOREF), in millivolts.
//
//	For example, use 3300 for a 3.3 V system or 5000 for a 5 V system.
//
// param offset The offset of the raw current sense signal for the Motoron
//
//	channel.  This is the same measurement that you would put into the
//	Motoron's "Current sense offset" variable using setCurrentSenseOffset(),
//	so see the documentation of that function for more info.
//	The offset is typically 10 for 5 V systems and 15 for 3.3 V systems,
//	(50*1024/referenceMv) but it can vary widely.
func (se *MXT550) CalculateCurrentLimit(milliAmps uint32, mType MotoronCurrentSenseType,
	referenceMv uint16, offset uint16) (limit uint16) {

	if milliAmps > 1000000 {
		milliAmps = 1000000
	}

	calc := (uint32(offset)*125+64)/128 + uint32(milliAmps)*20/(uint32(referenceMv)*(uint32(mType&3)))

	if calc > 1000 {
		limit = 1000
	} else {
		limit = uint16(calc)
	}

	return
}

// CurrentSenseUnitsMilliamps
//
// Calculates the units for the Motoron's current sense reading returned by
// getCurrentSenseProcessed(), in milliamps.
//
// To convert a reading from getCurrentSenseProcessed() to milliamps
// multiply it by the value returned from this function using 32-bit
// multiplication.  For example:
//
//	uint32_t ma = (uint32_t)processed * units;
//
// param type Specifies what type of Motoron you are using.
// param referenceMv The reference voltage (IOREF), in millivolts.
//
// For example, use 3300 for a 3.3 V system or 5000 for a 5 V system.
func (se *MXT550) CurrentSenseUnitsMilliamps(mType MotoronCurrentSenseType, referenceMv uint16) (milliAmps uint16) {

	calc := uint32(referenceMv) * (uint32(mType & 3)) * 25 / 512

	milliAmps = uint16(calc)
	return
}
