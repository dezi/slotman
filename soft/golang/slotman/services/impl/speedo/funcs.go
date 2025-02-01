package speedo

import (
	"errors"
	"fmt"
	"slotman/things/pololu/mxt550"
)

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
	}

	if motoron == nil {
		err = errors.New(fmt.Sprintf("motoron %d not found", track))
		return
	}

	speedValue := int16(800 * percent / 100)

	if speedValue < -800 {
		speedValue = -800
	}

	if speedValue > +800 {
		speedValue = +800
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
