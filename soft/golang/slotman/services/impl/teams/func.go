package teams

import (
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"path/filepath"
	"slotman/goodies/imaging"
	"slotman/services/type/slotman"
	"slotman/utils/log"
)

func (sv *Service) GetAllTeams() (teams []*slotman.Team) {
	teams = mockupTeams
	return
}

func GetScaledTeamLogo(logo string, size int) (img *image.RGBA, err error) {

	input, err := embedFs.Open(filepath.Join("embeds", logo))
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

func GetCircleTeamLogo(logo string, size int) (img *image.RGBA, err error) {

	input, err := embedFs.Open(filepath.Join("embeds", logo))
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

	img, err = imaging.ScaleToCircle(src, size, 2, "ffffff")
	return
}
