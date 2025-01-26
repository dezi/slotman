package turner

import (
	"slotman/services/impl/teams"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkDisplays()
	sv.displayTeams()
}

func (sv *Service) displayTeams() {

	sv.teamIndex = (sv.teamIndex + 1) % len(sv.teams)

	img, err := teams.GetScaledTeamLogo(sv.teams[sv.teamIndex].Logo, 240)
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
