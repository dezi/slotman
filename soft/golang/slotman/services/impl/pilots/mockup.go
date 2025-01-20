package pilots

import (
	"slotman/goodies/images"
	"slotman/services/type/slotman"
	"slotman/utils/simple"
)

var (
	mockups = []slotman.Pilot{
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Dennis",
			LastName:  "Zierahn",
			Team:      images.GetAllTeams()[0].Name,
			CarModel:  images.GetAllTeams()[0].Car,
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Patrick",
			LastName:  "Zierahn",
			Team:      images.GetAllTeams()[1].Name,
			CarModel:  images.GetAllTeams()[1].Car,
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Lukas",
			LastName:  "Zierahn",
			Team:      images.GetAllTeams()[2].Name,
			CarModel:  images.GetAllTeams()[2].Car,
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Kim",
			LastName:  "Zierahn",
			Team:      images.GetAllTeams()[3].Name,
			CarModel:  images.GetAllTeams()[3].Car,
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Susi",
			LastName:  "Brandt",
			Team:      images.GetAllTeams()[4].Name,
			CarModel:  images.GetAllTeams()[4].Car,
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Omar",
			LastName:  "Müller",
			Team:      images.GetAllTeams()[5].Name,
			CarModel:  images.GetAllTeams()[5].Car,
		},
	}
)
