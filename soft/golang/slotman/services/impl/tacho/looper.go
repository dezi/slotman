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

		if thisInputs == lastInputs {
			continue
		}

		sv.speedChan <- thisInputs
		lastInputs = thisInputs
	}
}

func (sv *Service) speedEval() {

	defer sv.waitGroup.Done()

	log.Printf("SpeedSensor eval started...")
	defer log.Printf("SpeedSensor eval stopped.")

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

				if now.UnixMilli()-state.time.UnixMilli() < 100 {
					continue
				}

				state.dirty = false

				if pin%2 == 0 {

					//
					// Speed measure pin.
					//

					log.Printf("############### speed pin=%02d track=%d active=%v",
						pin, pin>>1, state.active)
				}

				if pin%2 == 1 {

					//
					// Round measure pin.
					//

					if !state.active {
						state.round = state.round + 1
						log.Printf("############### round pin=%02d track=%d round=%d",
							pin, pin>>1, state.round)
					}
				}

				sv.speedStates[pin] = state
			}

		case inputs, ok := <-sv.speedChan:

			if !ok {
				return
			}

			now := time.Now()

			for pin := 0; pin < 16; pin++ {

				mask := uint16(1 << pin)
				active := inputs&mask != 0

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
