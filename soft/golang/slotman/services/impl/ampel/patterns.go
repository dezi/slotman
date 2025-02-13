package ampel

import (
	"slotman/services/type/slotman"
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"time"
)

func (sv *Service) patternRaceWaiting() {

	log.Printf("Pattern race waiting started...")
	defer log.Printf("Pattern race waiting done.")

	state := 0
	delay := 250

	for !sv.doExit && sv.ampelState == AmpelStateRaceWaiting {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			time.Sleep(time.Second)
			continue
		}

		sv.ampelLock.Lock()

		pins, _ := ampelGpio.ReadPins()
		pins &= 0x8000

		for track, trackState := range sv.waitingTracksReady {

			if track > 4 || trackState == slotman.TrackStateInactive {
				continue
			}

			if trackState == slotman.TrackStateActive {
				if state%2 == 0 {
					pins |= 1 << pinsYellow[4-track]
				}
			}

			if trackState == slotman.TrackStateReady {
				pins |= 1 << pinsYellow[4-track]
				pins |= 1 << pinsGreen[4-track]
			}
		}

		_ = ampelGpio.WritePins(pins)

		sv.ampelLock.Unlock()

		state = (state + 1) % 2

		wait := delay
		for !sv.doExit && sv.ampelState == AmpelStateRaceWaiting && wait > 0 {
			time.Sleep(time.Millisecond * 20)
			wait -= 20
		}
	}
}

func (sv *Service) patternRaceSuspended() {

	log.Printf("Pattern race suspended started...")
	defer log.Printf("Pattern race suspended done.")

	state := 0
	delay := 250

	for !sv.doExit && sv.ampelState == AmpelStateRaceSuspended {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			time.Sleep(time.Second)
			continue
		}

		sv.ampelLock.Lock()

		pins, _ := ampelGpio.ReadPins()
		pins &= 0x8000

		switch state {
		case 0: // * _ _ _ *
			pins |= 1 << pinsYellow[0]
			pins |= 1 << pinsYellow[4]
		case 1: // _ * _ * _
			pins |= 1 << pinsYellow[1]
			pins |= 1 << pinsYellow[3]
		case 2: // _ _ * _ _
			pins |= 1 << pinsYellow[2]
		case 3: // _ * _ * _
			pins |= 1 << pinsYellow[1]
			pins |= 1 << pinsYellow[3]
		case 4: // * _ _ _ *
			pins |= 1 << pinsYellow[0]
			pins |= 1 << pinsYellow[4]
		}

		_ = ampelGpio.WritePins(pins)

		sv.ampelLock.Unlock()

		state = (state + 1) % 4

		wait := delay
		for !sv.doExit && sv.ampelState == AmpelStateRaceSuspended && wait > 0 {
			time.Sleep(time.Millisecond * 20)
			wait -= 20
		}
	}
}

func (sv *Service) patternRaceRunning() {

	log.Printf("Pattern race running started...")
	defer log.Printf("Pattern race running done.")

	state := 0
	delay := 250

	for !sv.doExit && sv.ampelState == AmpelStateRaceRunning {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			time.Sleep(time.Second)
			continue
		}

		sv.ampelLock.Lock()

		pins, _ := ampelGpio.ReadPins()
		pins &= 0x8000

		switch state {
		case 0:
			pins |= 1 << pinsGreen[0]
			pins |= 1 << pinsGreen[1]
			pins |= 1 << pinsGreen[2]
			pins |= 1 << pinsGreen[3]
			pins |= 1 << pinsGreen[4]
		}

		_ = ampelGpio.WritePins(pins)

		sv.ampelLock.Unlock()

		state++

		if state > 0 {
			break
		}

		wait := delay
		for !sv.doExit && sv.ampelState == AmpelStateRaceRunning && wait > 0 {
			time.Sleep(time.Millisecond * 20)
			wait -= 20
		}
	}
}

func (sv *Service) patternRaceStart() {

	log.Printf("Pattern race start started...")
	defer log.Printf("Pattern race start done.")

	state := 0
	delay := 1000

	for !sv.doExit && sv.ampelState == AmpelStateRaceStart {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			time.Sleep(time.Second)
			continue
		}

		sv.ampelLock.Lock()

		pins, _ := ampelGpio.ReadPins()
		pins &= 0x8000

		//goland:noinspection GoDfaConstantCondition
		switch state {
		case 0:
			pins |= 0
		case 1:
			pins |= 1 << pinsRed[4]
		case 2:
			pins |= 1 << pinsRed[4]
			pins |= 1 << pinsRed[3]
		case 3:
			pins |= 1 << pinsRed[4]
			pins |= 1 << pinsRed[3]
			pins |= 1 << pinsRed[2]
		case 4:
			pins |= 1 << pinsRed[4]
			pins |= 1 << pinsRed[3]
			pins |= 1 << pinsRed[2]
			pins |= 1 << pinsRed[1]
		case 5:
			pins |= 1 << pinsRed[4]
			pins |= 1 << pinsRed[3]
			pins |= 1 << pinsRed[2]
			pins |= 1 << pinsRed[1]
			pins |= 1 << pinsRed[0]
		case 6:
			pins |= 1 << pinsGreen[0]
			pins |= 1 << pinsGreen[1]
			pins |= 1 << pinsGreen[2]
			pins |= 1 << pinsGreen[3]
			pins |= 1 << pinsGreen[4]
		}

		_ = ampelGpio.WritePins(pins)

		sv.ampelLock.Unlock()

		state++

		if state == 7 {
			break
		}

		wait := delay
		for !sv.doExit && sv.ampelState == AmpelStateRaceStart && wait > 0 {
			time.Sleep(time.Millisecond * 20)
			wait -= 20
		}
	}

	sv.rce.OnRaceStarted()
}

func (sv *Service) patternIdle() {

	log.Printf("Pattern idle started...")
	defer log.Printf("Pattern idle done.")

	state := 0
	delay := 500

	for !sv.doExit && sv.ampelState == AmpelStateIdle {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			time.Sleep(time.Second)
			continue
		}

		sv.ampelLock.Lock()

		pins, _ := ampelGpio.ReadPins()
		pins &= 0x8000

		switch state {
		case 0: // * _ _ _ *
			pins |= 1 << pinsGreen[0]
			pins |= 1 << pinsGreen[4]
		case 1: // _ * _ * _
			pins |= 1 << pinsGreen[1]
			pins |= 1 << pinsGreen[3]
		case 2: // _ _ * _ _
			pins |= 1 << pinsGreen[2]
		case 3: // _ * _ * _
			pins |= 1 << pinsGreen[1]
			pins |= 1 << pinsGreen[3]
		case 4: // * _ _ _ *
			pins |= 1 << pinsGreen[0]
			pins |= 1 << pinsGreen[4]
		}

		switch sv.roundsToGo {
		case 5:
			pins |= 1 << pinsYellow[4]
		case 10:
			pins |= 1 << pinsYellow[4]
			pins |= 1 << pinsYellow[3]
		case 25:
			pins |= 1 << pinsYellow[4]
			pins |= 1 << pinsYellow[3]
			pins |= 1 << pinsYellow[2]
		case 50:
			pins |= 1 << pinsYellow[4]
			pins |= 1 << pinsYellow[3]
			pins |= 1 << pinsYellow[2]
			pins |= 1 << pinsYellow[1]
		case 100:
			pins |= 1 << pinsYellow[4]
			pins |= 1 << pinsYellow[3]
			pins |= 1 << pinsYellow[2]
			pins |= 1 << pinsYellow[1]
			pins |= 1 << pinsYellow[0]
		}

		_ = ampelGpio.WritePins(pins)

		sv.ampelLock.Unlock()

		state = (state + 1) % 4

		wait := delay
		for !sv.doExit && sv.ampelState == AmpelStateIdle && wait > 0 {
			time.Sleep(time.Millisecond * 20)
			wait -= 20
		}
	}
}

func (sv *Service) patternTest() {

	for !sv.doExit {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			time.Sleep(time.Second)
			continue
		}

		for loop := 1; loop < 5; loop++ {

			_ = ampelGpio.WritePins(0xffff)
			//vals, _ := lightSensor.ReadPins()
			time.Sleep(time.Millisecond * 250)

			_ = ampelGpio.WritePins(0x0000)
			//vals, _ = lightSensor.ReadPins()
			time.Sleep(time.Millisecond * 250)
		}

		//_ = lightSensor.WritePin(5, mcp23017.PinLogicHi)
		//_ = lightSensor.WritePin(9, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(0, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(1, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(2, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(3, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(4, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(5, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(6, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(7, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(8, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(9, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(10, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(11, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(12, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(13, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(14, mcp23017.PinLogicHi)

		time.Sleep(time.Second * 5)
	}
}
