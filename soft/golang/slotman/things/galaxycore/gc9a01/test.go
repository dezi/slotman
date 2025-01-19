package gc9a01

import (
	"image/jpeg"
	"math/rand"
	"os"
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

	_ = gc9a01.SetFrame(Frame{X0: 0, Y0: 0, X1: ScreenWidth - 1, Y1: ScreenHeight - 1})

	log.Printf("Display GC8A01 test patterns.")

	_ = gc9a01.SetOrientation(2)

	f, err := os.Open("/home/liesa/dezi.profil.jpg")
	if err != nil {
		log.Cerror(err)
		return
	}

	rgbImage, err := jpeg.Decode(f)
	if err != nil {
		log.Cerror(err)
		return
	}

	_ = f.Close()

	log.Printf("Profil wid=%d hei=%d",
		rgbImage.Bounds().Size().X,
		rgbImage.Bounds().Size().Y)

	err = gc9a01.BlipFullImage(rgbImage)
	if err != nil {
		log.Cerror(err)
		return
	}

	time.Sleep(time.Second * 5)

	rawImage := make([]byte, ScreenWidth*ScreenHeight*3)

	for {

		color := make([]byte, 3)
		color[0] = byte(rand.Int31())
		color[1] = byte(rand.Int31())

		off := 0

		for x := 0; x < ScreenWidth; x++ {
			for y := 0; y < ScreenHeight; y++ {
				if x < y {
					color[2] = 0xFF
				} else {
					color[2] = 0x00
				}

				rawImage[off] = color[0]
				off++
				rawImage[off] = color[1]
				off++
				rawImage[off] = color[2]
				off++
			}
		}

		err = gc9a01.BlipFullImageRaw(rawImage)

		time.Sleep(time.Millisecond * 250)
	}
}
