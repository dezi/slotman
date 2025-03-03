package sc16is752

// i2cdetect -F 1
// i2cdetect -y 1

//
// https://github.com/nopnop2002/SC16IS752
//

//goland:noinspection GoUnusedGlobalVariable
var (
	ThingI2CAddresses = []byte{
		ThingI2CAddressAA, ThingI2CAddressAB,
		ThingI2CAddressBA, ThingI2CAddressBB,
	}
)

//goland:noinspection GoUnusedConst
const (
	//
	// Possible I2C addresses.
	//
	// A:VDD
	// B:GND
	// C:SCL
	// D:SDA
	//

	ThingI2CAddressAA byte = 0x90>>1 ^ 1
	ThingI2CAddressAB byte = 0x92>>1 ^ 1
	ThingI2CAddressAC byte = 0x94>>1 ^ 1
	ThingI2CAddressAD byte = 0x96>>1 ^ 1
	ThingI2CAddressBA byte = 0x98>>1 ^ 1
	ThingI2CAddressBB byte = 0x9A>>1 ^ 1
	ThingI2CAddressBC byte = 0x9C>>1 ^ 1
	ThingI2CAddressBD byte = 0x9E>>1 ^ 1
	ThingI2CAddressCA byte = 0xA0>>1 ^ 1
	ThingI2CAddressCB byte = 0xA2>>1 ^ 1
	ThingI2CAddressCC byte = 0xA4>>1 ^ 1
	ThingI2CAddressCD byte = 0xA6>>1 ^ 1
	ThingI2CAddressDA byte = 0xA8>>1 ^ 1
	ThingI2CAddressDB byte = 0xAA>>1 ^ 1
	ThingI2CAddressDC byte = 0xAC>>1 ^ 1
	ThingI2CAddressDD byte = 0xAE>>1 ^ 1

	//
	// Channel select
	//

	ChannelA byte = 0x00
	ChannelB byte = 0x01

	//
	// General Registers
	//

	RegRHR       byte = 0x00
	RegTHR       byte = 0x00
	RegIER       byte = 0x01
	RegFCR       byte = 0x02
	RegIIR       byte = 0x02
	RegLCR       byte = 0x03
	RegMCR       byte = 0x04
	RegLSR       byte = 0x05
	RegMSR       byte = 0x06
	RegSPR       byte = 0x07
	RegTCR       byte = 0x06
	RegTLR       byte = 0x07
	RegTxLvl     byte = 0x08
	RegRxLvl     byte = 0x09
	RegIoDir     byte = 0x0A
	RegIoState   byte = 0x0B
	RegIoIntEna  byte = 0x0C
	RegIoControl byte = 0x0E
	RegEfCr      byte = 0x0F

	//
	// Special Registers
	//

	RegDLL byte = 0x00
	RegDLH byte = 0x01

	//
	// Enhanced Registers
	//

	RegEfr   byte = 0x02
	RegXon1  byte = 0x04
	RegXon2  byte = 0x05
	RegXoff1 byte = 0x06
	RegXoff2 byte = 0x07

	ParityNone   byte = 0
	ParityOdd    byte = 1
	ParityEven   byte = 2
	ParityForce1 byte = 3
	ParityForce0 byte = 4
)
