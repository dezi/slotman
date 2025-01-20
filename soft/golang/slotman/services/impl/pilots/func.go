package pilots

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"path/filepath"
	"slotman/utils/log"
	"strings"
)

func GetScaledPilotImage(pilot string, size int) (img *image.RGBA, err error) {

	file := fmt.Sprintf("profile-%s.jpg", strings.ToLower(pilot))

	input, err := embedFs.Open(filepath.Join("embeds", file))
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
