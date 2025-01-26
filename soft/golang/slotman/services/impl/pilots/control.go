package pilots

import (
	"math/rand"
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

	var err error

	allTeams := sv.tms.GetAllTeams()
	teamIndex := rand.Int() % len(allTeams)

	for _, mockupPilot := range mockupPilots {

		teamIndex = (teamIndex + 1) % len(allTeams)

		mockupPilot.Team = allTeams[teamIndex].Name
		mockupPilot.Car = allTeams[teamIndex].Car

		mockupPilot.ProfilePic, err = sv.loadMockupPilotProfile(
			mockupPilot.FirstName,
			mockupPilot.LastName)
		log.Cerror(err)

		sv.UpdatePilot(mockupPilot)
	}
}
