package bmp280

// i2cdetect -F 1
// i2cdetect -y 1

type Register byte
type PowerMode byte
type PowerInterval byte
type Oversampling byte
type IrrFilter byte

//goland:noinspection GoUnusedConst
const (
	ThingI2CAddress = 0x58

	RegisterId        Register = 0xd0
	RegisterReset     Register = 0xe0
	RegisterStatus    Register = 0xf3
	RegisterCtrlMeas  Register = 0xf4
	RegisterConfig    Register = 0xf5
	RegisterPressMsb  Register = 0xf7
	RegisterPressLsb  Register = 0xf8
	RegisterPressXlsb Register = 0xf9
	RegisterTempMsb   Register = 0xfa
	RegisterTempLsb   Register = 0xfb
	RegisterTempXlsb  Register = 0xfc

	PowerModeSleep  PowerMode = 0x00
	PowerModeForced PowerMode = 0x01
	PowerModeNormal PowerMode = 0x03

	PowerInterval1ms    PowerInterval = 0x00
	PowerInterval65ms   PowerInterval = 0x01
	PowerInterval125ms  PowerInterval = 0x02
	PowerInterval250ms  PowerInterval = 0x03
	PowerInterval500ms  PowerInterval = 0x04
	PowerInterval1000ms PowerInterval = 0x05
	PowerInterval2000ms PowerInterval = 0x06
	PowerInterval4000ms PowerInterval = 0x07

	OversamplingDisable Oversampling = 0x00
	Oversampling1       Oversampling = 0x01
	Oversampling2       Oversampling = 0x02
	Oversampling4       Oversampling = 0x03
	Oversampling8       Oversampling = 0x04
	Oversampling16      Oversampling = 0x07

	IrrFilterOff IrrFilter = 0x00
	IrrFilter2   IrrFilter = 0x02
	IrrFilter4   IrrFilter = 0x03
	IrrFilter8   IrrFilter = 0x04
	IrrFilter16  IrrFilter = 0x05
)

var (
	compensationRegs = []byte{0x88, 0x8A, 0x8C, 0x8E, 0x90, 0x92, 0x94, 0x96, 0x98, 0x9A, 0x9C, 0x9E}
)
