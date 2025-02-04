package tacho

import (
	"slotman/utils/log"
	"time"
)

func (sv *Service) speedRead() {

	defer sv.waitGroup.Done()

	log.Printf("SpeedSensor read started...")
	defer log.Printf("SpeedSensor read stopped.")

	var err error

	thisInputs := uint16(0x0000)
	lastInputs := uint16(0xffff)

	for !sv.doExit {

		speedSensor := sv.speedSensor
		if speedSensor == nil {
			break
		}

		thisInputs, err = speedSensor.ReadPins()
		if err != nil {
			log.Cerror(err)
			time.Sleep(time.Millisecond * 100)
			continue
		}

		//
		// Remark: missing pull down resistors...
		//

		thisInputs &= 0x000f

		if thisInputs == lastInputs {
			continue
		}

		sv.speedChan <- SpeedRead{
			pinStates: thisInputs,
			readTime:  time.Now(),
		}

		lastInputs = thisInputs
	}
}

func (sv *Service) speedEval() {

	defer sv.waitGroup.Done()

	log.Printf("SpeedSensor eval started...")
	defer log.Printf("SpeedSensor eval stopped.")

	var active bool

	for !sv.doExit {

		speedSensor := sv.speedSensor
		if speedSensor == nil {
			break
		}

		select {

		case <-time.After(time.Millisecond):

			now := time.Now()

			for pin := 0; pin < 16; pin++ {

				state := sv.speedStates[pin]

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

				sv.speedStates[pin] = state

				//
				// Todo push message here if proxy.
				//

				sv.handleLocalSpeed(pin, state)
			}

		case speedRead, ok := <-sv.speedChan:

			if !ok {
				return
			}

			now := speedRead.readTime
			inputs := speedRead.pinStates

			for pin := 0; pin < 16; pin++ {

				mask := uint16(1 << pin)
				active = inputs&mask != 0

				state := sv.speedStates[pin]

				if state.active != active {
					state.active = active
					state.dirty = true
					state.time = now
					sv.speedStates[pin] = state
				}
			}
		}
	}
}

func (sv *Service) handleLocalSpeed(pin int, state SpeedState) {

	track := pin >> 1
	active := state.active

	if pin%2 == 0 {

		//
		// Start + speed measure pin.
		//

		//log.Printf("Speed pin=%02d track=%d active=%v",
		//	pin, track, active)

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

			state2 := sv.speedStates[pin+1]

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

		log.Printf("Round pin=%02d track=%d active=%v", pin, track, active)
	}
}
