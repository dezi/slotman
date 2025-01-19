package gc9a01

import (
	"slotman/utils/log"
)

var gc9a01 *GC9A01

func TestDisplay() {

	log.Printf("GC8A01 display test started...")

	gc9a01 = NewGC9A01("/dev/spidev0.0", 25)

	err := gc9a01.OpenSensor()
	if err != nil {
		return
	}

	log.Printf("GC8A01 device SPI0-0 opened.")

}
