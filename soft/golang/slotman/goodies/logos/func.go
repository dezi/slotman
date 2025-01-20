package logos

import (
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"slotman/utils/log"
)

func GetAllTeams() (teams []Team) {
	teams = allTeams
	return
}

func GetScaledTeamLogo(logo string, size int) (img *image.RGBA, err error) {

	input, err := embedFs.Open(logo)
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() { _ = input.Close() }()

	src, err := png.Decode(input)
	if err != nil {
		log.Cerror(err)
		return
	}

	img = image.NewRGBA(image.Rect(0, 0, size, size))
	draw.NearestNeighbor.Scale(img, img.Rect, src, src.Bounds(), draw.Over, nil)

	return
}
