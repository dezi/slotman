package gc9a01

//
// https://www.waveshare.com/wiki/1.28inch_LCD_Module
// https://www.az-delivery.de/en/products/1-28-zoll-rundes-tft-display
// https://dronebotworkshop.com/gc9a01/
//
// https://github.com/carlfriess/GC9A01_demo
// https://github.com/adafruit/Adafruit_GC9A01A
//

type Command byte

type Orientation byte
type ColorMode byte

//goland:noinspection GoUnusedConst
const (
	CommandColAddrSet Command = 0x2A
	CommandRowAddrSet Command = 0x2B

	CommandOrientation Command     = 0x36
	Orientation0       Orientation = 0x18
	Orientation90      Orientation = 0x28
	Orientation180     Orientation = 0x48
	Orientation270     Orientation = 0x88

	CommandColorMode Command   = 0x3A
	ColorMode12Bit   ColorMode = 0x03
	ColorMode16Bit   ColorMode = 0x05
	ColorMode18Bit   ColorMode = 0x06

	CommandMemWr     Command = 0x2C
	CommandMemWrCont Command = 0x3C
)

const (
	ScreenWidth  = 240
	ScreenHeight = 240
)
