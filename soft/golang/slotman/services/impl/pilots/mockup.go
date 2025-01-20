package pilots

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"slotman/services/type/slotman"
	"slotman/utils/log"
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

func (sv *Service) loadMockupPilotProfile(pilot string) (base64Img string, err error) {

	file := fmt.Sprintf("profile-%s.jpg", strings.ToLower(pilot))

	data, err := embedFs.ReadFile(filepath.Join("embeds", file))
	if err != nil {
		log.Cerror(err)
		return
	}

	base64Img = "data:image/jpeg;base64, " + base64.StdEncoding.EncodeToString([]byte(data))
	return
}
