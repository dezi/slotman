package turner

import (
	"image"
	"image/png"
	"path/filepath"
	"slotman/utils/imaging"
	"slotman/utils/log"
)

func (sv *Service) getSlotmanLogo() (rgba *image.RGBA, err error) {

	input, err := embedFs.Open(filepath.Join("embeds", "slotman-logo.png"))
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

	rgba, err = imaging.ScaleToCircle(src, 240, 0, "")
	return
}

func (sv *Service) blipFullImage(img *image.RGBA) (err error) {

	if sv.isProxyClient {

		req := &Turner{
			Area:      AreaTurner,
			What:      TurnerWhatBlipFull,
			BlipImage: imaging.GetImageRawData(img),
		}

		_, err = sv.prx.ProxyRequest(req)
		log.Cerror(err)

	} else {

		if sv.turnDisplay1 != nil {
			_ = sv.turnDisplay1.Initialize()
			_ = sv.turnDisplay1.BlipFullImage(img)
		}

		if sv.turnDisplay2 != nil {
			_ = sv.turnDisplay2.Initialize()
			_ = sv.turnDisplay2.BlipFullImage(img)
		}
	}

	return
}
