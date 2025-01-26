package ld2461

//
// https://github.com/Chreece/LD2461-ESPHome
// https://drive.google.com/drive/folders/14_KgZpL4Th2LTRuq_W0X_ePjHu_Z6lcS
//
// Motion target detection and tracking
//

var (
	baudRates = []int{9600}
)

// ralph 0176 72198103

const (
	commandSetBaudrate    = 0x01
	commandSetReporting   = 0x02
	commandGetReporting   = 0x03
	commandSetRegions     = 0x04
	commandDisableRegions = 0x05
	commandGetRegions     = 0x06
	commandGetCoordinates = 0x07
	commandGetNumTargets  = 0x08
	commandReadFirmware   = 0x09
	commandRestoreFactory = 0x0A
)
