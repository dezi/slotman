package turner

import (
	"image"
	"slotman/utils/imaging"
	"slotman/utils/log"
)

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
