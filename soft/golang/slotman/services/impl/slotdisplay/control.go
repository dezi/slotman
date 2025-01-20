package slotdisplay

import (
	"github.com/fogleman/gg"
	"image"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkDisplays()
}

func (sv *Service) checkDisplays() {

	if sv.turnDisplay1 == nil {
		sv.turnDisplay1 = gc9a01.NewGC9A01("/dev/spidev0.0", 25)
		_ = sv.turnDisplay1.OpenSensor()
	}

	if sv.turnDisplay2 == nil {
		sv.turnDisplay2 = gc9a01.NewGC9A01("/dev/spidev0.1", 25)
		_ = sv.turnDisplay2.OpenSensor()
	}

	//
	// Re-initialize every 60 seconds.
	//

	_ = sv.turnDisplay1.Initialize()
	_ = sv.turnDisplay2.Initialize()

	img, err := sv.turnDisplay1.LoadScaledImage("/home/liesa/dezi.profil.jpg")
	if err != nil {
		return
	}

	log.Printf("Profil wid=%d hei=%d", img.Bounds().Size().X, img.Bounds().Size().Y)

	dc := gg.NewContextForRGBA(img.(*image.RGBA))
	dc.SetRGB255(0xff, 0xff, 0xff)
	dc.SetFontFace(sv.faceRegularNormal)
	dc.DrawStringAnchored("P. Zierahn", 120, 200, 0.5, 0.0)

	err = sv.turnDisplay1.BlipFullImage(img)
	if err != nil {
		log.Cerror(err)
		return
	}
}
