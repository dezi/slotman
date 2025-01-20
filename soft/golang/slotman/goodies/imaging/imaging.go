package imaging

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/draw"
	"image"
)

func ScaleToCircle(src image.Image, size, borderWidth int, borderColor string) (circle *image.RGBA, err error) {

	srcScaled := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.NearestNeighbor.Scale(srcScaled, srcScaled.Rect, src, src.Bounds(), draw.Over, nil)

	circle = image.NewRGBA(image.Rect(0, 0, size, size))

	dc := gg.NewContextForRGBA(circle)

	dc.SetHexColor("ff000000")
	dc.Fill()

	dc.SetHexColor(borderColor)
	dc.DrawRoundedRectangle(0, 0, float64(size), float64(size), float64(size/2))
	dc.Fill()

	size -= borderWidth

	dc.DrawRoundedRectangle(0, 0, float64(size), float64(size), float64(size/2))
	dc.Clip()

	dc.DrawImage(srcScaled, 0, 0)

	return
}
