package pilots

import (
	"slotman/services/type/slotman"
	"slotman/utils/simple"
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
