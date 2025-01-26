package turner

import (
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkDisplays()
	sv.displayTeams()
	sv.displayPilots()
	sv.loopCount++
}

func (sv *Service) displayPilots() {

	if sv.loopCount%2 == 0 {
		return
	}

	pilots := sv.plt.GetAllPilots()

	sv.pilotIndex = (sv.pilotIndex + 1) % len(pilots)

	img, err := sv.plt.GetScaledPilotPic(pilots[sv.pilotIndex], 240)
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

func (sv *Service) displayTeams() {

	if sv.loopCount%2 == 1 {
		return
	}

	teams := sv.tms.GetAllTeams()

	sv.teamIndex = (sv.teamIndex + 1) % len(teams)

	img, err := sv.tms.GetScaledTeamLogo(teams[sv.teamIndex], 240)
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
