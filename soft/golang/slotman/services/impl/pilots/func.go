package pilots

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"path/filepath"
	"slotman/goodies/imaging"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func (sv *Service) UpdatePilot(pilot *slotman.Pilot) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	sv.pilots[pilot.AppUuid] = pilot

	sv.pilotProfileFull[pilot.AppUuid] = nil
	sv.pilotProfileSmall[pilot.AppUuid] = nil

	if pilot.ProfilePic != "" {

		img, err := decodeBaseImage(pilot.ProfilePic)
		if err != nil {

			log.Cerror(err)

		} else {

			var full *image.RGBA
			full, err = imaging.ScaleToCircle(img, 240, 0, "")
			if err == nil {
				sv.pilotProfileFull[pilot.AppUuid] = full
			}

			var small *image.RGBA
			small, err = imaging.ScaleToCircle(img, 40, 2, "ff0000")
			if err == nil {
				sv.pilotProfileSmall[pilot.AppUuid] = small
			}
		}
	}

	if pilot.IsMockup {
		return
	}

	//
	// System starts up with mockups.
	// Remove all mockup pilots when
	// the first real pilot registers.
	//

	var deletes []simple.UUIDHex

	for appUuid, pilotRec := range sv.pilots {
		if pilotRec.IsMockup {
			deletes = append(deletes, appUuid)
		}
	}

	for _, appUuid := range deletes {
		delete(sv.pilots, appUuid)
		delete(sv.pilotProfileFull, appUuid)
		delete(sv.pilotProfileSmall, appUuid)
	}
}

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
