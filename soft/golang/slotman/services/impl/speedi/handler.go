package speedi

import (
	"encoding/json"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

func (sv *Service) speedControlHandler(track int) {

	var lastTime int64
	var rawSpeed uint16

	for !sv.doExit {

		select {

		case <-time.After(time.Millisecond * 50):
			continue

		case rawSpeed = <-sv.speedControlChannels[track]:
		}

		if sv.isProxyServer {
			err := sv.pushLocalSpeed(track, rawSpeed, &lastTime)
			log.Cerror(err)
		} else {
			err := sv.handleLocalSpeed(track, rawSpeed, &lastTime)
			log.Cerror(err)
		}
	}
}

func (sv *Service) pushLocalSpeed(track int, rawSpeed uint16, lastTime *int64) (err error) {

	speedi := &Speedi{
		Uuid:     simple.NewUuidHex(),
		Area:     AreaSpeedi,
		What:     SpeediWhatSpeed,
		Track:    track,
		RawSpeed: rawSpeed,
		Ok:       true,
		Err:      "",
	}

	speediBytes, err := json.Marshal(speedi)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.prx.ProxyBroadcast(speediBytes)
	log.Cerror(err)

	if rawSpeed > 7000 || (lastTime == nil || time.Now().Unix()-*lastTime > 5) {
		log.Printf("Speed track=%d rawSpeed=%d", track, rawSpeed)
		*lastTime = time.Now().Unix()
	}

	return
}

func (sv *Service) handleLocalSpeed(track int, rawSpeed uint16, lastTime *int64) (err error) {

	if rawSpeed == 0 {
		sv.speedControlAttached[track] = false
		return
	}

	sv.speedControlAttached[track] = true

	pcc := &PlayerControlCurve{
		BrkPercent: 5,
		MinPercent: 25,
		MaxPercent: 70,
	}

	spc := sv.speedControlCalibrations[track]

	speed := 100 * float64(rawSpeed-spc.MinValue) / float64(spc.MaxValue-spc.MinValue)

	speedPcc := speed
	if speed <= 0.1 {
		speedPcc = pcc.BrkPercent
	} else {
		speedPcc = pcc.MinPercent + speed*(pcc.MaxPercent-pcc.MinPercent)/100
	}

	_ = sv.sdo.SetSpeed(track, speedPcc, false)

	if speed != 0 && (lastTime == nil || time.Now().Unix()-*lastTime > 5) {
		log.Printf("Speed track=%d speedPcc=%5.1f speed=%5.1f rawSpeed=%d", track, speedPcc, speed, rawSpeed)
		if lastTime != nil {
			*lastTime = time.Now().Unix()
		}
	}

	return
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
