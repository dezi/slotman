package pilots

import (
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

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
