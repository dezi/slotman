package pilots

import (
	"slotman/services/type/slotman"
	"slotman/utils/simple"
)

var (
	mockup = []slotman.Pilot{
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Dennis",
			LastName:  "Zierahn",
			Team:      "",
			CarModel:  "",
		},
	}
)
