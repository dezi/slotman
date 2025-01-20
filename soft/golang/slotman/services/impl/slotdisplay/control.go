package slotdisplay

import (
	"slotman/goodies/logos"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkDisplays()
	sv.displayTeam()
}

func (sv *Service) displayTeam() {

	img, err := logos.GetScaledTeamLogo(sv.teamDefs[sv.teamIndex].Logo, 240)
	if err != nil {
		log.Cerror(err)
		return
	}

	_ = sv.turnDisplay1.Initialize()
	_ = sv.turnDisplay1.BlipFullImage(img)

	_ = sv.turnDisplay2.Initialize()
	_ = sv.turnDisplay2.BlipFullImage(img)
}

func (sv *Service) checkDisplays() {

	if sv.turnDisplay1 == nil {
		sv.turnDisplay1 = gc9a01.NewGC9A01("/dev/spidev0.0", 25)
		_ = sv.turnDisplay1.Open()
	}

	if sv.turnDisplay2 == nil {
		sv.turnDisplay2 = gc9a01.NewGC9A01("/dev/spidev0.1", 25)
		_ = sv.turnDisplay2.Open()
	}

	//
	// Re-initialize every 60 seconds.
	//

	//img, err := sv.turnDisplay1.LoadScaledImage("/home/liesa/dezi.profil.jpg")
	//if err != nil {
	//	return
	//}
	//
	//log.Printf("Profil wid=%d hei=%d", img.Bounds().Size().X, img.Bounds().Size().Y)
	//
	//dc := gg.NewContextForRGBA(img.(*image.RGBA))
	//dc.SetRGB255(0xff, 0xff, 0xff)
	//dc.SetFontFace(sv.faceRegularNormal)
	//dc.DrawStringAnchored("P. Zierahn", 120, 200, 0.5, 0.0)
	//
	//err = sv.turnDisplay1.BlipFullImage(img)
	//if err != nil {
	//	log.Cerror(err)
	//	return
	//}
}
