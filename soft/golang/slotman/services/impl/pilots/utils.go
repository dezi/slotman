package pilots

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"
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
