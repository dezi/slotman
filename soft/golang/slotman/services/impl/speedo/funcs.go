package speedo

import (
	"errors"
	"fmt"
	"slotman/services/type/slotman"
	"slotman/things/pololu/mxt550"
	"slotman/utils/log"
)

func (sv *Service) SetTrackEnable(track int, enable bool) {

	if track < 0 || track >= slotman.MaxTracks {
		return
	}

	log.Printf("SetTrackEnable track=%d enable=%v", track, enable)

	sv.tracksEnable[track] = enable
}

func (sv *Service) SetTrackEnableAll(enable bool) {

	log.Printf("SetTrackEnableAll enable=%v", enable)

	for track := 0; track < slotman.MaxTracks; track++ {
		sv.tracksEnable[track] = enable
	}
}

func (sv *Service) GetMotoronsAttached() (tracks []bool) {

	tracks = make([]bool, 8)

	if sv.mxt550Motoron1 != nil {
		tracks[0] = true
		tracks[1] = true
	}

	if sv.mxt550Motoron2 != nil {
		tracks[2] = true
		tracks[3] = true
	}

	if sv.mxt550Motoron3 != nil {
		tracks[4] = true
		tracks[5] = true
	}

	if sv.mxt550Motoron4 != nil {
		tracks[6] = true
		tracks[7] = true
	}

	return
}

func (sv *Service) SetSpeed(track int, percent float64, now bool) (err error) {

	var motor byte
	var motoron *mxt550.MXT550

	switch track {
	case 0:
		motor = 1
		motoron = sv.mxt550Motoron1
	case 1:
		motor = 2
		motoron = sv.mxt550Motoron1
	case 2:
		motor = 1
		motoron = sv.mxt550Motoron2
	case 3:
		motor = 2
		motoron = sv.mxt550Motoron2
	case 4:
		motor = 1
		motoron = sv.mxt550Motoron3
	case 5:
		motor = 2
		motoron = sv.mxt550Motoron3
	case 6:
		motor = 1
		motoron = sv.mxt550Motoron4
	case 7:
		motor = 2
		motoron = sv.mxt550Motoron4
	}

	if motoron == nil {
		err = errors.New(fmt.Sprintf("motoron %d not found", track))
		return
	}

	speedValue := int16(0)

	if sv.tracksEnable[track] {

		speedValue = int16(800 * percent / 100)

		if speedValue < -800 {
			speedValue = -800
		}

		if speedValue > +800 {
			speedValue = +800
		}
	}

	if now {
		err = motoron.SetSpeedNow(motor, speedValue)
	} else {
		err = motoron.SetSpeed(motor, speedValue)
	}

	if err != nil {
		err = errors.New(fmt.Sprintf("%s addr=%02x", err.Error(), motoron.GetThingAddress()))
	}

	return
}
