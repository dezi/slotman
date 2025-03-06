package ld2451

//
// Motion target detection and tracking
//

var (
	baudRates = []int{115200, 9600}

	validBaudRates = map[int]int{
		9600:   0x0001,
		19200:  0x0002,
		38400:  0x0003,
		57600:  0x0004,
		115200: 0x0005,
		230400: 0x0006,
		256000: 0x0007,
		460800: 0x0008,
	}
)

type DetectionMode byte

const (
	commandSetDetection   byte = 0x02
	commandGetDetection   byte = 0x12
	commandSetSensitivity byte = 0x03
	commandGetSensitivity byte = 0x13

	commandEnableConfigurations byte = 0xff
	commandEndConfiguration     byte = 0xfe

	commandReadFirmwareVersion   byte = 0xa0
	commandSetBaudrate           byte = 0xa1
	commandRestoreFactorySetting byte = 0xa2
	commandRestartModule         byte = 0xa3

	DetectionModeOnlyAway     DetectionMode = 0x00
	DetectionModeOnlyApproach DetectionMode = 0x01
	DetectionModeBoth         DetectionMode = 0x02
)
