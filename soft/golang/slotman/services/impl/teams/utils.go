package teams

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"
)

func (sv *Service) loadMockupTeamLogo(teamName string) (base64Img string, err error) {

	lowName := strings.ToLower(teamName)
	lowName = strings.ReplaceAll(lowName, " ", "-")

	file := fmt.Sprintf("logo-%s.png", lowName)

	data, err := embedFs.ReadFile(filepath.Join("embeds", file))
	if err != nil {
		return
	}

	base64Img = "data:image/png;base64, " + base64.StdEncoding.EncodeToString(data)
	return
}

func (sv *Service) loadMockupCarPic(teamName string) (base64Img string, err error) {

	lowName := strings.ToLower(teamName)
	lowName = strings.ReplaceAll(lowName, " ", "-")

	file := fmt.Sprintf("car-%s.png", lowName)

	data, err := embedFs.ReadFile(filepath.Join("embeds", file))
	if err != nil {
		return
	}

	base64Img = "data:image/png;base64, " + base64.StdEncoding.EncodeToString(data)
	return
}
