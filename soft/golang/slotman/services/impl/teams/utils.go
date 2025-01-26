package teams

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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

func decodeBaseImage(base64Image string) (img image.Image, err error) {

	base64Png := "data:image/png;base64, "
	if strings.HasPrefix(base64Image, base64Png) {

		base64Data := strings.TrimPrefix(base64Image, base64Png)

		var data []byte
		data, err = base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			return
		}

		img, err = png.Decode(strings.NewReader(string(data)))
		return
	}

	base64Jpeg := "data:image/jpeg;base64, "
	if strings.HasPrefix(base64Image, base64Jpeg) {

		base64Data := strings.TrimPrefix(base64Image, base64Jpeg)

		var data []byte
		data, err = base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			return
		}

		img, err = jpeg.Decode(strings.NewReader(string(data)))
		return
	}

	err = errors.New("unsupported image format")
	return
}
