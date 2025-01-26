package pilots

import (
	"math/rand"
	"slotman/services/impl/teams"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.loadMockups()
}

func (sv *Service) loadMockups() {

	if len(sv.pilots) > 0 {
		return
	}

	log.Printf("Loading pilot mockups start...")
	defer log.Printf("Loading pilot mockups done.")

	allTeams := teams.GetAllTeams()
	teamIndex := rand.Int() % len(allTeams)

	for _, mockupPilot := range mockupPilots {

		teamIndex = (teamIndex + 1) % len(allTeams)

		mockupPilot.Team = allTeams[teamIndex].Name
		mockupPilot.Car = allTeams[teamIndex].Car

		mockupPilot.ProfilePic, _ = sv.loadMockupPilotProfile(
			mockupPilot.FirstName,
			mockupPilot.LastName)

		sv.UpdatePilot(mockupPilot)
	}
}
