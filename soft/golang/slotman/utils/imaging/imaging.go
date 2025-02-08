package imaging

import (
	"encoding/base64"
	"errors"
	"github.com/fogleman/gg"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

func DecodeBase64Image(base64Image string) (img image.Image, err error) {

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

func GetImageRawData(rgba *image.RGBA) (raw []byte) {

	wid := rgba.Bounds().Size().X
	hei := rgba.Bounds().Size().Y
	pix := rgba.Pix

	raw = make([]byte, wid*hei*3)

	src := 0
	dst := 0

	for x := 0; x < wid; x++ {

		stride := src

		for y := 0; y < hei; y++ {

			raw[dst] = pix[stride]
			stride++
			dst++

			raw[dst] = pix[stride]
			stride++
			dst++

			raw[dst] = pix[stride]
			stride++
			dst++

			stride++
		}

		src += rgba.Stride
	}

	return
}
