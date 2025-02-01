package speedi

import (
	"slotman/things"
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

		//_ = sv.SetSpeed(track, speedPcc, false)

		if speed != 0 || time.Now().Unix()-lastTime.Unix() > 5 {
			log.Printf("Speed track=%d speedPcc=%5.1f speed=%5.1f rawSpeed=%d", track, speedPcc, speed, rawSpeed)
			lastTime = time.Now()
		}
	}
}

func (sv *Service) OnThingOpened(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing opened uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnThingClosed(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing closed uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnThingStarted(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing started uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnThingStopped(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing stopped uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnADConversion(thing things.Thing, input int, value uint16) {

	track := -1

	address := thing.GetThingAddress()

	if address == speedControl1Addr {
		track = 0 + input
	}

	if address == speedControl2Addr {
		track = 4 + input
	}

	if track >= maxTracks {
		return
	}

	//
	// Make sure we spend minimum time in this handler.
	//

	sv.speedControlChannels[track] <- value
}
