package turner

import (
	"slotman/goodies/teamdefs"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkDisplays()
	sv.displayTeams()
}

func (sv *Service) displayTeams() {

	sv.teamIndex = (sv.teamIndex + 1) % len(sv.teamDefs)

	img, err := teamdefs.GetScaledTeamLogo(sv.teamDefs[sv.teamIndex].Logo, 240)
	if err != nil {
		log.Cerror(err)
		return
	}

	if sv.turnDisplay1 != nil {
		_ = sv.turnDisplay1.Initialize()
		_ = sv.turnDisplay1.BlipFullImage(img)
	}

	if sv.turnDisplay2 != nil {
		_ = sv.turnDisplay2.Initialize()
		_ = sv.turnDisplay2.BlipFullImage(img)
	}
}

func (sv *Service) checkDisplays() {

	if sv.turnDisplay1 == nil {
		turnDisplay1 := gc9a01.NewGC9A01("/dev/spidev0.0", 25)
		tryErr := turnDisplay1.Open()
		if tryErr == nil {
			sv.turnDisplay1 = turnDisplay1
		}
	}

	if sv.turnDisplay2 == nil {
		turnDisplay2 := gc9a01.NewGC9A01("/dev/spidev0.1", 25)
		tryErr := turnDisplay2.Open()
		if tryErr == nil {
			sv.turnDisplay2 = turnDisplay2
		}
	}

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
