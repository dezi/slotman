package gc9a01

//
// https://www.waveshare.com/wiki/1.28inch_LCD_Module
// https://www.az-delivery.de/en/products/1-28-zoll-rundes-tft-display
// https://dronebotworkshop.com/gc9a01/
//
// https://github.com/carlfriess/GC9A01_demo
//

const (
	COL_ADDR_SET       byte = 0x2A
	ROW_ADDR_SET       byte = 0x2B
	MEM_WR             byte = 0x2C
	COLOR_MODE         byte = 0x3A
	COLOR_MODE__12_BIT byte = 0x03
	COLOR_MODE__16_BIT byte = 0x05
	COLOR_MODE__18_BIT byte = 0x06
	MEM_WR_CONT        byte = 0x3C
)
