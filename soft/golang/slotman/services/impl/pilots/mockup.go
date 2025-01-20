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
			AppUuid:    simple.NewUuidHex(),
			FirstName:  "Dennis",
			LastName:   "Zierahn",
			ProfilePic: "",
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
			FirstName: "Lewis",
			LastName:  "Hamilton",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Max",
			LastName:  "Verstappen",
		},
		{
			AppUuid:   simple.NewUuidHex(),
			FirstName: "Fernando",
			LastName:  "Alonso",
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

	base64Img = "data:image/jpeg;base64, " + base64.StdEncoding.EncodeToString([]byte(data))
	return
}
