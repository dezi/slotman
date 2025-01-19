package gc9a01

import (
	"math/rand"
	"slotman/utils/log"
	"time"
)

var gc9a01 *GC9A01

func TestDisplay() {

	log.Printf("Display GC8A01 test started...")

	gc9a01 = NewGC9A01("/dev/spidev0.0", 25)

	err := gc9a01.OpenSensor()
	if err != nil {
		return
	}

	log.Printf("Display GC8A01 device SPI0-0 opened.")

	_ = gc9a01.SetFrame(Frame{
		X0: 0,
		Y0: 0,
		X1: 239,
		Y1: 239,
	})

	log.Printf("Display GC8A01 test patterns.")

	chunk := 4

	line := make([]byte, screenWidth*3*chunk)

	for {

		color := make([]byte, 3)
		color[0] = byte(rand.Int31())
		color[1] = byte(rand.Int31())

		//for x := 0; x < screenWidth; x++ {
		//	for y := 0; y < screenHeight; y++ {
		//		if x < y {
		//			color[2] = 0xFF
		//		} else {
		//			color[2] = 0x00
		//		}
		//		if x == 0 && y == 0 {
		//			_ = gc9a01.writeMem(color)
		//		} else {
		//			_ = gc9a01.writeMemCont(color)
		//		}
		//	}
		//}

		off := 0
		first := true

		for x := 0; x < screenWidth; x++ {

			if x%chunk == 0 {
				off = 0
			}

			for y := 0; y < screenHeight; y++ {
				if x < y {
					color[2] = 0xFF
				} else {
					color[2] = 0x00
				}

				line[off] = color[0]
				off++
				line[off] = color[1]
				off++
				line[off] = color[2]
				off++
			}

			if (x+1)%chunk == 0 {
				if first {
					_ = gc9a01.writeMem(line)
					first = false
				} else {
					_ = gc9a01.writeMemCont(line)
				}
			}
		}

		//_ = gc9a01.writeMem(line)

		time.Sleep(time.Millisecond * 250)
	}
}
