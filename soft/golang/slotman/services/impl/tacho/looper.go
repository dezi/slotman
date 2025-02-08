package tacho

import (
	"encoding/json"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

func (sv *Service) tachoRead() {

	defer sv.waitGroup.Done()

	log.Printf("TachoSensor read started...")
	defer log.Printf("TachoSensor read stopped.")

	var err error

	thisInputs := uint16(0x0000)
	lastInputs := uint16(0xffff)

	for !sv.doExit {

		tachoSensor := sv.tachoSensor
		if tachoSensor == nil {
			break
		}

		thisInputs, err = tachoSensor.ReadPins()
		if err != nil {
			//log.Cerror(err)
			//time.Sleep(time.Millisecond * 100)
			continue
		}

		//
		// Remark: missing pull down resistors...
		//

		thisInputs &= 0x000f

		if thisInputs == lastInputs {
			continue
		}

		log.Printf("############ thisInputs=%04x", thisInputs)

		sv.tachoChan <- TachoRead{
			pinStates: thisInputs,
			readTime:  time.Now(),
		}

		lastInputs = thisInputs
	}
}

func (sv *Service) tachoEval() {

	defer sv.waitGroup.Done()

	log.Printf("TachoSensor eval started...")
	defer log.Printf("TachoSensor eval stopped.")

	var active bool

	for !sv.doExit {

		tachoSensor := sv.tachoSensor
		if tachoSensor == nil {
			break
		}

		select {

		case <-time.After(time.Millisecond):

			now := time.Now()

			for pin := 0; pin < 16; pin++ {

				state := sv.tachoStates[pin]

				if !state.dirty {
					continue
				}

				//
				// Pin state needs to be stable for
				// 100 ms to be processed.
				//

				if now.UnixMilli()-state.time.UnixMilli() < 100 {
					continue
				}

				state.dirty = false

				sv.tachoStates[pin] = state

				if sv.isProxyServer {
					sv.pushLocalTacho(pin, state)
				} else {
					sv.handleLocalTacho(pin, state)
				}
			}

		case tachoRead, ok := <-sv.tachoChan:

			if !ok {
				return
			}

			now := tachoRead.readTime
			inputs := tachoRead.pinStates

			for pin := 0; pin < 16; pin++ {

				mask := uint16(1 << pin)
				active = inputs&mask != 0

				state := sv.tachoStates[pin]

				if state.active != active {
					state.active = active
					state.dirty = true
					state.time = now
					sv.tachoStates[pin] = state
				}
			}
		}
	}
}

func (sv *Service) pushLocalTacho(pin int, state TachoState) {

	track := pin >> 1
	active := state.active

	log.Printf("Tacho push pin=%02d track=%d active=%v", pin, track, active)

	tacho := &Tacho{
		Uuid:   simple.NewUuidHex(),
		Area:   AreaTacho,
		What:   TachoWhatTacho,
		Pin:    pin,
		Active: active,
		Time:   state.time,
		Ok:     true,
		Err:    "",
	}

	tachoBytes, err := json.Marshal(tacho)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.prx.ProxyBroadcast(tachoBytes)
	log.Cerror(err)
}

func (sv *Service) handleLocalTacho(pin int, state TachoState) {

	track := pin >> 1
	active := state.active

	log.Printf("Tacho local pin=%02d track=%d active=%v", pin, track, active)

	if pin%2 == 0 {

		//
		// Start + speed measure pin.
		//

		sv.mapsLock.Lock()
		trackState := sv.trackStates[track]
		sv.mapsLock.Unlock()

		if active {

			if !trackState.IsAtStart {

				//
				// Enter start position.
				//

				go sv.OnEnterStartPosition(track)
			}

		} else {

			if trackState.IsAtStart {

				//
				// Leave start position.
				//

				go sv.OnLeaveStartPosition(track)
			}

			//
			// Take speed if possible.
			//

			state2 := sv.tachoStates[pin+1]

			microSecs := state.time.UnixMicro() - state2.time.UnixMicro()

			if microSecs > 0 && microSecs < 1000000 {

				//
				// Time measurement possible.
				//

				speed := float64(sensorDistMM) / float64(microSecs)
				// mm / sec
				speed *= 1000000
				// mm / h
				speed *= 3600
				// cm / h
				speed /= 10
				// m / h
				speed /= 100
				// km / h
				speed /= 1000
				// Scale 43 / 2
				speed *= 43 / 2

				go sv.OnSpeedMeasurement(track, speed)

				go sv.OnRoundCompleted(track)
			}
		}
	}

	if pin%2 == 1 {
		//
		// Speed measure pin 2.
		// The state is only maintained
		// for further processing.
		//
	}
}
