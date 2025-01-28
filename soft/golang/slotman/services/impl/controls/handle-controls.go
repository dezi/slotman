package controls

import (
	"slotman/utils/log"
	"time"
)

func (sv *Service) speedControlHandler(track int) {

	pcc := &PlayerControlCurve{
		BrkPercent: 5,
		MinPercent: 25,
		MaxPercent: 70,
	}

	spc := sv.speedControlCalibrations[track]

	var rawSpeed uint16
	var lastTime time.Time

	for !sv.doExit {

		select {

		case <-time.After(time.Millisecond * 50):
			continue

		case rawSpeed = <-sv.speedControlChannels[track]:
		}

		if rawSpeed == 0 {
			sv.speedControlAttached[track] = false
			continue
		}

		sv.speedControlAttached[track] = true

		speed := 100 * float64(rawSpeed-spc.MinValue) / float64(spc.MaxValue-spc.MinValue)

		speedPcc := speed
		if speed <= 0.1 {
			speedPcc = pcc.BrkPercent
		} else {
			speedPcc = pcc.MinPercent + speed*(pcc.MaxPercent-pcc.MinPercent)/100
		}

		err := sv.SetSpeed(track, speedPcc, false)
		log.Cerror(err)

		if speed != 0 || time.Now().Unix()-lastTime.Unix() > 5 {
			log.Printf("Speed track=%d speedPcc=%5.1f speed=%5.1f rawSpeed=%d", track, speedPcc, speed, rawSpeed)
			lastTime = time.Now()
		}
	}
}
