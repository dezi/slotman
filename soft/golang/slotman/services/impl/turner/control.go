package turner

import (
	"fmt"
	"github.com/fogleman/gg"
	"slotman/services/type/race"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkDisplays()
	sv.displayState()
	//sv.displayTeams()
	//sv.displayPilots()
	sv.loopCount++
}

func (sv *Service) displayState() {

	if sv.isProxyServer {
		return
	}

	if sv.loopCount%3 != 0 {
		return
	}

	img, err := sv.getSlotmanLogo()
	if err != nil {
		log.Cerror(err)
		return
	}

	raceState := sv.rce.GetRaceState()

	motoronsAttached := sv.sdo.GetMotoronsAttached()
	speedControls := sv.sdi.GetSpeedControlsAttached()

	if raceState == race.RaceStateIdle {

	}

	dc := gg.NewContextForRGBA(img)

	dc.DrawRectangle(0, 0, 240, 240)
	dc.SetHexColor("#00000080")
	dc.Fill()

	dc.SetHexColor("e0bf78")
	dc.SetFontFace(sv.faceBoldLarge)
	dc.DrawStringAnchored("Hardware", 120, 44, 0.5, 0.0)
	dc.DrawStringAnchored("________", 120, 48, 0.5, 0.0)

	for inx := 0; inx < 4; inx++ {

		dc.SetHexColor("e0bf78")
		text := fmt.Sprintf("%d:", inx+1)
		dc.DrawString(text, 60, float64(86+inx*36))

		if motoronsAttached[inx] {
			dc.SetHexColor("00ff00")
		} else {
			dc.SetHexColor("ff0000")
		}
		dc.DrawString("M", 100, float64(86+inx*36))

		if speedControls[inx] {
			dc.SetHexColor("00ff00")
		} else {
			dc.SetHexColor("ff0000")
		}
		dc.DrawString("C", 130, float64(86+inx*36))

		if speedControls[inx] {
			dc.SetHexColor("00ff00")
		} else {
			dc.SetHexColor("ff0000")
		}
		dc.DrawString("T", 160, float64(86+inx*36))
	}

	err = sv.blipFullImage(img)
	log.Cerror(err)
}

func (sv *Service) displayPilots() {

	if sv.isProxyServer {
		return
	}

	if sv.loopCount%3 != 1 {
		return
	}

	pilots := sv.plt.GetAllPilots()

	sv.pilotIndex = (sv.pilotIndex + 1) % len(pilots)

	img, err := sv.plt.GetScaledPilotPic(pilots[sv.pilotIndex], 240)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.blipFullImage(img)
	log.Cerror(err)
}

func (sv *Service) displayTeams() {

	if sv.isProxyServer {
		return
	}

	if sv.loopCount%3 != 2 {
		return
	}

	teams := sv.tms.GetAllTeams()

	sv.teamIndex = (sv.teamIndex + 1) % len(teams)

	img, err := sv.tms.GetScaledTeamLogo(teams[sv.teamIndex], 240)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.blipFullImage(img)
	log.Cerror(err)
}

func (sv *Service) checkDisplays() {

	if sv.isProxyClient {
		return
	}

	if sv.turnDisplay1 == nil {

		turnDisplay1 := gc9a01.NewGC9A01("/dev/spidev0.0", 25)
		turnDisplay1.SetHandler(sv)

		tryErr := turnDisplay1.Open()
		if tryErr == nil {
			sv.turnDisplay1 = turnDisplay1
		} else {
			log.Cerror(tryErr)
		}
	}

	if sv.turnDisplay2 == nil {

		turnDisplay2 := gc9a01.NewGC9A01("/dev/spidev0.1", 25)
		turnDisplay2.SetHandler(sv)

		tryErr := turnDisplay2.Open()
		if tryErr == nil {
			sv.turnDisplay2 = turnDisplay2
		} else {
			log.Cerror(tryErr)
		}
	}
}
