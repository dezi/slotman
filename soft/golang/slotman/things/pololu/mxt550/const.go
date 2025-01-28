package mxt550

//
// https://www.pololu.com/product/5079
// https://github.com/pololu/motoron-arduino
// https://www.berrybase.de/pololu-motoron-m2t550-dual-i2c-motor-controller-steuert-2-gleichstrommotoren-18-22v-16a
//
// https://www.pololu.com/docs/0J84/6
// https://www.pololu.com/docs/0J84/9
//

//goland:noinspection GoUnusedConst
var (
	//
	// CustomAddresses
	// ThingAddresses
	//
	// 0x0f => Motoron with JMP1 shortened to GND.
	// 0x10 => Motoron with JMP1 default open.
	//
	// 0x18-0x1f => Motoron with custom address.
	//
	// Probing will scan custom addresses first.
	//
	// A motoron with address 0x10 is attached unmodified.
	//
	// A motoron with address 0x0f will be assigned to
	// the first unoccupied custom address and then reset.
	//
	CustomAddresses = []byte{0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}
	ThingAddresses  = append(CustomAddresses, 0x0f, 0x10)
)

type MotoronCmd byte
type MotoronProtocolOption byte
type MotoronProtocolOptions byte
type MotoronVar byte
type MotoronMvar byte
type MotoronStatusFlag byte
type MotoronStatusFlags uint16
type MotoronSetting byte
type MotoronProductId uint16

//goland:noinspection GoUnusedConst
const (
	MotoronCmdGetFirmwareVersion          MotoronCmd = 0x87
	MotoronCmdSetProtocolOptions          MotoronCmd = 0x8B
	MotoronCmdReadEeprom                  MotoronCmd = 0x93
	MotoronCmdWriteEeprom                 MotoronCmd = 0x95
	MotoronCmdReinitialize                MotoronCmd = 0x96
	MotoronCmdReset                       MotoronCmd = 0x99
	MotoronCmdGetVariables                MotoronCmd = 0x9A
	MotoronCmdSetVariable                 MotoronCmd = 0x9C
	MotoronCmdCoastNow                    MotoronCmd = 0xA5
	MotoronCmdClearMotorFault             MotoronCmd = 0xA6
	MotoronCmdClearLatchedStatusFlags     MotoronCmd = 0xA9
	MotoronCmdSetLatchedStatusFlags       MotoronCmd = 0xAC
	MotoronCmdSetBraking                  MotoronCmd = 0xB1
	MotoronCmdSetBrakingNow               MotoronCmd = 0xB2
	MotoronCmdSetSpeed                    MotoronCmd = 0xD1
	MotoronCmdSetSpeedNow                 MotoronCmd = 0xD2
	MotoronCmdSetBufferedSpeed            MotoronCmd = 0xD4
	MotoronCmdSetAllSpeeds                MotoronCmd = 0xE1
	MotoronCmdSetAllSpeedsNow             MotoronCmd = 0xE2
	MotoronCmdSetAllBufferedSpeeds        MotoronCmd = 0xE4
	MotoronCmdSetAllSpeedsUsingBuffers    MotoronCmd = 0xF0
	MotoronCmdSetAllSpeedsNowUsingBuffers MotoronCmd = 0xF3
	MotoronCmdResetCommandTimeout         MotoronCmd = 0xF5
	MotoronCmdMultiDeviceErrorCheck       MotoronCmd = 0xF9
	MotoronCmdMultiDeviceWrite            MotoronCmd = 0xFA

	MotoronSettingFactoryResetCode        MotoronSetting = 0
	MotoronSettingDeviceNumber            MotoronSetting = 1
	MotoronSettingAlternativeDeviceNumber MotoronSetting = 3
	MotoronSettingCommunicationOptions    MotoronSetting = 5
	MotoronSettingBaudDivider             MotoronSetting = 6
	MotoronSettingResponseDelay           MotoronSetting = 8

	MotoronVarProtocolOptions MotoronVar = 0
	MotoronVarStatusFlags     MotoronVar = 1
	MotoronVarVinVoltage      MotoronVar = 3
	MotoronVarCommandTimeout  MotoronVar = 5
	MotoronVarErrorResponse   MotoronVar = 7
	MotoronVarErrorMask       MotoronVar = 8
	MotoronVarJumperState     MotoronVar = 10
	MotoronVarUartFaults      MotoronVar = 11

	MotoronMvarPwmMode                     MotoronVar = 1
	MotoronMvarTargetSpeed                 MotoronVar = 2
	MotoronMvarTargetBrakeAmount           MotoronVar = 4
	MotoronMvarCurrentSpeed                MotoronVar = 6
	MotoronMvarBufferedSpeed               MotoronVar = 8
	MotoronMvarMaxAccelForward             MotoronVar = 10
	MotoronMvarMaxAccelReverse             MotoronVar = 12
	MotoronMvarMaxDecelForward             MotoronVar = 14
	MotoronMvarMaxDecelReverse             MotoronVar = 16
	MotoronMvarStartingSpeedForward        MotoronVar = 18
	MotoronMvarStartingSpeedReverse        MotoronVar = 20
	MotoronMvarDirectionChangeDelayForward MotoronVar = 22
	MotoronMvarDirectionChangeDelayReverse MotoronVar = 23
	MotoronMvarMaxDecelTemporary           MotoronVar = 24
	MotoronMvarCurrentLimit                MotoronVar = 26
	MotoronMvarCurrentSenseRaw             MotoronVar = 28
	MotoronMvarCurrentSenseSpeed           MotoronVar = 30
	MotoronMvarCurrentSenseProcessed       MotoronVar = 32
	MotoronMvarCurrentSenseOffset          MotoronVar = 34
	MotoronMvarCurrentSenseMinimumDivisor  MotoronVar = 35

	MotoronProtocolOptionCrcForCommands  MotoronProtocolOption = 0
	MotoronProtocolOptionCrcForResponses MotoronProtocolOption = 1
	MotoronProtocolOptionI2cGeneralCall  MotoronProtocolOption = 2

	DefaultProtocolOptions MotoronProtocolOptions = (1 << MotoronProtocolOptionCrcForCommands) |
		(1 << MotoronProtocolOptionCrcForResponses) |
		(1 << MotoronProtocolOptionI2cGeneralCall)

	MotoronCommunicationOption7bitResponses     = 0
	MotoronCommunicationOption14bitDeviceNumber = 1
	MotoronCommunicationOptionErrIsDe           = 2

	MotoronStatusFlagProtocolError         MotoronStatusFlag = 0
	MotoronStatusFlagCrcError              MotoronStatusFlag = 1
	MotoronStatusFlagCommandTimeoutLatched MotoronStatusFlag = 2
	MotoronStatusFlagMotorFaultLatched     MotoronStatusFlag = 3
	MotoronStatusFlagNoPowerLatched        MotoronStatusFlag = 4
	MotoronStatusFlagUartError             MotoronStatusFlag = 5
	MotoronStatusFlagReset                 MotoronStatusFlag = 9
	MotoronStatusFlagCommandTimeout        MotoronStatusFlag = 10
	MotoronStatusFlagMotorFaulting         MotoronStatusFlag = 11
	MotoronStatusFlagNoPower               MotoronStatusFlag = 12
	MotoronStatusFlagErrorActive           MotoronStatusFlag = 13
	MotoronStatusFlagMotorOutputEnabled    MotoronStatusFlag = 14
	MotoronStatusFlagMotorDriving          MotoronStatusFlag = 15

	DefaultErrorMask MotoronStatusFlags = (1 << MotoronStatusFlagCommandTimeout) |
		(1 << MotoronStatusFlagReset)

	MotoronUartFaultFraming         = 0
	MotoronUartFaultNoise           = 1
	MotoronUartFaultHardwareOverrun = 2
	MotoronUartFaultSoftwareOverrun = 3

	MotoronErrorResponseCoast    = 0
	MotoronErrorResponseBrake    = 1
	MotoronErrorResponseCoastNow = 2
	MotoronErrorResponseBrakeNow = 3

	MotoronPwmModeDefault = 0
	MotoronPwmMode1Khz    = 1
	MotoronPwmMode2Khz    = 2
	MotoronPwmMode4Khz    = 3
	MotoronPwmMode5Khz    = 4
	MotoronPwmMode10Khz   = 5
	MotoronPwmMode20Khz   = 6
	MotoronPwmMode40Khz   = 7
	MotoronPwmMode80Khz   = 8

	MotoronJmp1Installed    = 0
	MotoronJmp1NotInstalled = 1

	MotoronClearMotorFaultUnconditional = 0

	MotoronErrorCheckContinue byte = 0x3C
	MotoronErrorCheckDone     byte = 0x00

	MotoronMaxSpeed                = 800
	MotoronMaxAccel                = 6400
	MotoronMaxDirectionChangeDelay = 250

	MotoronLatchedStatusFlags = 0x03FF
	MotoronMaxErrorMask       = 0x07FF

	MotoronMaxCommandTimeout = 16250

	MotoronMinBaudRate uint32 = 245
	MotoronMaxBaudRate uint32 = 1000000

	MotoronProductIdM3S256 MotoronProductId = 0x00CC
	MotoronProductIdM3H256 MotoronProductId = 0x00CC
	MotoronProductIdM2S    MotoronProductId = 0x00CD
	MotoronProductIdM2H    MotoronProductId = 0x00CD
	MotoronProductIdM2T256 MotoronProductId = 0x00CE
	MotoronProductIdM2U256 MotoronProductId = 0x00CF
	MotoronProductIdM1T256 MotoronProductId = 0x00D0
	MotoronProductIdM1U256 MotoronProductId = 0x00D1
	MotoronProductIdM3S550 MotoronProductId = 0x00D2
	MotoronProductIdM3H550 MotoronProductId = 0x00D2
	MotoronProductIdM2T550 MotoronProductId = 0x00D3
	MotoronProductIdM2U550 MotoronProductId = 0x00D4
	MotoronProductIdM1T550 MotoronProductId = 0x00D5
	MotoronProductIdM1U550 MotoronProductId = 0x00D6
)

// MotoronCurrentSenseType
// Specifies what type of Motoron is being used, for the purposes of current
// limit and current sense calculations.
type MotoronCurrentSenseType byte

//goland:noinspection GoUnusedConst
const (
	Motoron18v18 MotoronCurrentSenseType = 0b0001
	Motoron24v14 MotoronCurrentSenseType = 0b0101
	Motoron18v20 MotoronCurrentSenseType = 0b1010
	Motoron24v16 MotoronCurrentSenseType = 0b1101
)

// MotoronVinSenseType
// Specifies what type of Motoron is being used, for the purposes of
// converting raw VIN voltage readings to millivolts.
type MotoronVinSenseType byte

//goland:noinspection GoUnusedConst
const (
	Motoron256 MotoronVinSenseType = 0b0000 // M*256 Motorons
	MotoronHp  MotoronVinSenseType = 0b0010 // High-power Motorons
	Motoron550 MotoronVinSenseType = 0b0011 // M*550 Motorons
)

type MotoronCurrentSenseReading struct {
	Raw       uint16
	Speed     int16
	Processed uint16
}

var (
	MotoronCmd2Str = map[MotoronCmd]string{
		MotoronCmdGetFirmwareVersion:          "GetFirmwareVersion",
		MotoronCmdSetProtocolOptions:          "SetProtocolOptions",
		MotoronCmdReadEeprom:                  "ReadEeprom",
		MotoronCmdWriteEeprom:                 "WriteEeprom",
		MotoronCmdReinitialize:                "Reinitialize",
		MotoronCmdReset:                       "Reset",
		MotoronCmdGetVariables:                "GetVariables",
		MotoronCmdSetVariable:                 "SetVariable",
		MotoronCmdCoastNow:                    "CoastNow",
		MotoronCmdClearMotorFault:             "ClearMotorFault",
		MotoronCmdClearLatchedStatusFlags:     "ClearLatchedStatusFlags",
		MotoronCmdSetLatchedStatusFlags:       "SetLatchedStatusFlags",
		MotoronCmdSetBraking:                  "SetBraking",
		MotoronCmdSetBrakingNow:               "BrakingNow",
		MotoronCmdSetSpeed:                    "SetSpeed",
		MotoronCmdSetSpeedNow:                 "SetSpeed",
		MotoronCmdSetBufferedSpeed:            "SetBufferedSpeed",
		MotoronCmdSetAllSpeeds:                "SetAllSpeeds",
		MotoronCmdSetAllSpeedsNow:             "SetAllSpeedsNow",
		MotoronCmdSetAllBufferedSpeeds:        "SetAllBufferedSpeeds",
		MotoronCmdSetAllSpeedsUsingBuffers:    "SetAllSpeedsUsingBuffers",
		MotoronCmdSetAllSpeedsNowUsingBuffers: "SetAllSpeedsNowUsingBuffers",
		MotoronCmdResetCommandTimeout:         "ResetCommandTimeout",
		MotoronCmdMultiDeviceErrorCheck:       "MultiDeviceErrorCheck",
		MotoronCmdMultiDeviceWrite:            "MultiDeviceWrite",
	}

	MotoronVar2Str = map[MotoronVar]string{
		MotoronVarProtocolOptions: "ProtocolOptions",
		MotoronVarStatusFlags:     "StatusFlags",
		MotoronVarVinVoltage:      "VinVoltage",
		MotoronVarCommandTimeout:  "CommandTimeout",
		MotoronVarErrorResponse:   "ErrorResponse",
		MotoronVarErrorMask:       "ErrorMask",
		MotoronVarJumperState:     "JumperState",
		MotoronVarUartFaults:      "UartFaults",
	}

	MotoronMvar2Str = map[MotoronVar]string{

		MotoronMvarPwmMode:                     "PwmMode",
		MotoronMvarTargetSpeed:                 "TargetSpeed",
		MotoronMvarTargetBrakeAmount:           "TargetBrakeAmount",
		MotoronMvarCurrentSpeed:                "CurrentSpeed",
		MotoronMvarBufferedSpeed:               "BufferedSpeed",
		MotoronMvarMaxAccelForward:             "MaxAccelForward",
		MotoronMvarMaxAccelReverse:             "MaxAccelReverse",
		MotoronMvarMaxDecelForward:             "MaxDecelForward",
		MotoronMvarMaxDecelReverse:             "MaxDecelReverse",
		MotoronMvarStartingSpeedForward:        "StartingSpeedForward",
		MotoronMvarStartingSpeedReverse:        "StartingSpeedReverse",
		MotoronMvarDirectionChangeDelayForward: "DirectionChangeDelayForward",
		MotoronMvarDirectionChangeDelayReverse: "DirectionChangeDelayReverse",
		MotoronMvarMaxDecelTemporary:           "MaxDecelTemporary",
		MotoronMvarCurrentLimit:                "CurrentLimit",
		MotoronMvarCurrentSenseRaw:             "CurrentSenseRaw",
		MotoronMvarCurrentSenseSpeed:           "CurrentSenseSpeed",
		MotoronMvarCurrentSenseProcessed:       "CurrentSenseProcessed",
		MotoronMvarCurrentSenseOffset:          "CurrentSenseOffset",
		MotoronMvarCurrentSenseMinimumDivisor:  "CurrentSenseMinimumDivisor",
	}

	MotoronStatusFlag2Str = map[MotoronStatusFlag]string{
		MotoronStatusFlagProtocolError:         "ProtocolError",
		MotoronStatusFlagCrcError:              "CrcError",
		MotoronStatusFlagCommandTimeoutLatched: "CommandTimeoutLatched",
		MotoronStatusFlagMotorFaultLatched:     "MotorFaultLatched",
		MotoronStatusFlagNoPowerLatched:        "NoPowerLatched",
		MotoronStatusFlagUartError:             "UartError",
		MotoronStatusFlagReset:                 "Reset",
		MotoronStatusFlagCommandTimeout:        "CommandTimeout",
		MotoronStatusFlagMotorFaulting:         "MotorFaulting",
		MotoronStatusFlagNoPower:               "NoPower",
		MotoronStatusFlagErrorActive:           "ErrorActive",
		MotoronStatusFlagMotorOutputEnabled:    "MotorOutputEnabled",
		MotoronStatusFlagMotorDriving:          "MotorDriving",
	}
)
