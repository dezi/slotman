package pilots

import (
	"math/rand"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

var (
	mockupPilots = []*slotman.Pilot{
		{
			FirstName: "Dennis",
			LastName:  "Zierahn",
		},
		{
			FirstName: "Patrick",
			LastName:  "Zierahn",
		},
		{
			FirstName: "Lukas",
			LastName:  "Zierahn",
		},
		{
			FirstName: "Kim",
			LastName:  "Zierahn",
		},
		{
			FirstName: "Susi",
			LastName:  "Brandt",
		},
		{
			FirstName: "Omar",
			LastName:  "Müller",
		},
		{
			FirstName: "Alex",
			LastName:  "Albon",
			Team:      "Williams-Martini",
		},
		{
			FirstName: "Carlos",
			LastName:  "Sainz",
			Team:      "Ferrari",
		},
		{
			FirstName: "Charles",
			LastName:  "Leclerc",
			Team:      "Ferrari",
		},
		{
			FirstName: "Esteban",
			LastName:  "Ocon",
			Team:      "Alpine",
		},
		{
			FirstName: "Fernando",
			LastName:  "Alonso",
			Team:      "Aston Martin",
		},
		{
			FirstName: "George",
			LastName:  "Russell",
			Team:      "Mercedes-AMG",
		},
		{
			FirstName: "Lance",
			LastName:  "Stroll",
			Team:      "Aston Martin",
		},
		{
			FirstName: "Lando",
			LastName:  "Norris",
			Team:      "McLaren",
		},
		{
			FirstName: "Lewis",
			LastName:  "Hamilton",
			Team:      "Mercedes-AMG",
		},
		{
			FirstName: "Max",
			LastName:  "Verstappen",
			Team:      "Red Bull-Oracle",
		},
		{
			FirstName: "Nico",
			LastName:  "Hulkenberg",
			Team:      "Haas",
		},
		{
			FirstName: "Oliver",
			LastName:  "Bearman",
			Team:      "Ferrari",
		},
		{
			FirstName: "Oscar",
			LastName:  "Piastri",
			Team:      "McLaren",
		},
		{
			FirstName: "Pierre",
			LastName:  "Gasly",
			Team:      "Alpine",
		},
		{
			FirstName: "Yuki",
			LastName:  "Tsunoda",
			Team:      "Racing Bulls",
		},
	}
)

func (sv *Service) loadMockups() {

	if len(sv.pilots) > 0 {
		return
	}

	log.Printf("Loading pilot mockups start...")
	defer log.Printf("Loading pilot mockups done.")

	var err error
	var team *slotman.Team

	allTeams := sv.tms.GetAllTeams()
	teamIndex := rand.Int() % len(allTeams)

	for _, mp := range mockupPilots {

		mp.Uuid = simple.UuidHexFromSha256([]byte(mp.FirstName + "|" + mp.LastName))

		teamIndex = (teamIndex + 1) % len(allTeams)

		if mp.Team == "" {

			mp.Team = allTeams[teamIndex].Name
			mp.Car = allTeams[teamIndex].Car

		} else {

			team, err = sv.tms.GetTeam(mp.Team)
			if err != nil {
				log.Cerror(err)
				continue
			}

			mp.Team = team.Name
			mp.Car = team.Car
		}

		mp.ProfilePic, err = sv.loadMockupPilotProfile(
			mp.FirstName,
			mp.LastName)
		log.Cerror(err)

		sv.UpdatePilot(mp)
	}
}
