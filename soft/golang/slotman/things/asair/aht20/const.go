package aht20

// i2cdetect -F 1
// i2cdetect -y 1

type Register byte

//goland:noinspection GoUnusedConst
const (
	ThingI2CAddress = 0x38

	RegisterInit    Register = 0xbe
	RegisterReset   Register = 0xba
	RegisterMeasure Register = 0xac
)
