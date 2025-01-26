package pilots

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"slotman/services/type/slotman"
	"slotman/utils/simple"
	"strings"
)

var (
	mockupPilots = []*slotman.Pilot{
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Dennis",
			LastName:  "Zierahn",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Patrick",
			LastName:  "Zierahn",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Lukas",
			LastName:  "Zierahn",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Kim",
			LastName:  "Zierahn",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Susi",
			LastName:  "Brandt",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Omar",
			LastName:  "Müller",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Alex",
			LastName:  "Albon",
			Team:      "Williams-Martini",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Carlos",
			LastName:  "Sainz",
			Team:      "Ferrari",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Charles",
			LastName:  "Leclerc",
			Team:      "Ferrari",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Esteban",
			LastName:  "Ocon",
			Team:      "Alpine",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Fernando",
			LastName:  "Alonso",
			Team:      "Aston Martin",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "George",
			LastName:  "Russell",
			Team:      "Mercedes-AMG",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Lance",
			LastName:  "Stroll",
			Team:      "Aston Martin",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Lando",
			LastName:  "Norris",
			Team:      "McLaren",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Lewis",
			LastName:  "Hamilton",
			Team:      "Mercedes-AMG",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Max",
			LastName:  "Verstappen",
			Team:      "Red Bull-Oracle",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Nico",
			LastName:  "Hulkenberg",
			Team:      "Haas",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Oliver",
			LastName:  "Bearman",
			Team:      "Ferrari",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Oscar",
			LastName:  "Piastri",
			Team:      "McLaren",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Pierre",
			LastName:  "Gasly",
			Team:      "Alpine",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Yuki",
			LastName:  "Tsunoda",
			Team:      "Racing Bulls",
		},
	}
)

func (sv *Service) loadMockupPilotProfile(pilotFirstname, pilotLastName string) (base64Img string, err error) {

	lowFirst := strings.ToLower(pilotFirstname)
	lowLast := strings.ToLower(pilotLastName)

	file := fmt.Sprintf("profile-%s-%s.jpg", lowFirst, lowLast)

	data, err := embedFs.ReadFile(filepath.Join("embeds", file))
	if err != nil {

		file = fmt.Sprintf("profile-%s.jpg", lowFirst)

		data, err = embedFs.ReadFile(filepath.Join("embeds", file))
		if err != nil {
			return
		}
	}

	base64Img = "data:image/jpeg;base64, " + base64.StdEncoding.EncodeToString(data)
	return
}
