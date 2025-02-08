package ampel

import (
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"time"
)

func (sv *Service) patternIdle() {

	log.Printf("Pattern idle started...")
	defer log.Printf("Pattern idle done.")

	state := 0
	delay := 500

	for !sv.doExit && sv.ampelState == AmpelStateIdle {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			break
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
			break
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
