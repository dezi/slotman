package pilots

import (
	"errors"
	"golang.org/x/image/draw"
	"image"
	"slotman/services/type/slotman"
	"slotman/utils/imaging"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sort"
)

func (sv *Service) GetPilot(pilotUuid simple.UUIDHex) (pilot *slotman.Pilot, err error) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	pilot = sv.pilots[pilotUuid]

	if pilot == nil {
		err = errors.New("no pilot for uuid")
		log.Cerror(err)
		return
	}

	return
}

func (sv *Service) GetAllPilots() (pilots []*slotman.Pilot) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	for _, pilot := range sv.pilots {
		pilots = append(pilots, pilot)
	}

	sort.Slice(pilots, func(i, j int) bool {

		if pilots[i].LastName != pilots[j].LastName {
			return pilots[i].LastName < pilots[j].LastName
		}

		return pilots[i].FirstName < pilots[j].FirstName
	})

	return
}

func (sv *Service) GetScaledPilotPic(pilot *slotman.Pilot, size int) (img *image.RGBA, err error) {

	src, err := imaging.DecodeBase64Image(pilot.ProfilePic)
	if err != nil {
		log.Cerror(err)
		return
	}

	img = image.NewRGBA(image.Rect(0, 0, size, size))
	draw.NearestNeighbor.Scale(img, img.Rect, src, src.Bounds(), draw.Over, nil)

	return
}

func (sv *Service) GetCircularPilotPic(pilot *slotman.Pilot, size int) (img *image.RGBA, err error) {

	img, err = sv.GetScaledPilotPic(pilot, size)
	if err != nil {
		return
	}

	img, err = imaging.ScaleToCircle(img, size, 8, "e0bf78")
	return
}

func (sv *Service) UpdatePilot(pilot *slotman.Pilot) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	sv.pilots[pilot.Uuid] = pilot

	sv.pilotProfileFull[pilot.Uuid] = nil
	sv.pilotProfileSmall[pilot.Uuid] = nil

	if pilot.ProfilePic != "" {

		img, err := imaging.DecodeBase64Image(pilot.ProfilePic)
		if err != nil {

			log.Cerror(err)

		} else {

			var full *image.RGBA
			full, err = imaging.ScaleToCircle(img, 240, 0, "")
			if err == nil {
				sv.pilotProfileFull[pilot.Uuid] = full
			}

			var small *image.RGBA
			small, err = imaging.ScaleToCircle(img, 40, 2, "ff0000")
			if err == nil {
				sv.pilotProfileSmall[pilot.Uuid] = small
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
