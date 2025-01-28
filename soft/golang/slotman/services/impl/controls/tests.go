package controls

import (
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"time"
)

type speedState struct {
	active bool
	dirty  bool
	round  int
	time   time.Time
}

var speedSensor *mcp23017.MCP23017
var speedChan = make(chan uint16, 10)
var speedStates = make(map[int]speedState)

func (sv *Service) speedTest(sensor *mcp23017.MCP23017) {

	if speedSensor != nil {
		return
	}

	speedSensor = sensor

	err := speedSensor.SetPinDirections(0x000f)
	if err != nil {
		log.Cerror(err)
		return
	}

	go sv.speedRead()
	go sv.speedEval()
}

func (sv *Service) speedRead() {

	log.Printf("############### speedSensor read started...")
	defer log.Printf("############### speedSensor read stopped.")

	var err error

	thisInputs := uint16(0x0000)
	lastInputs := uint16(0xffff)

	for speedSensor != nil {

		thisInputs, err = speedSensor.ReadPins()
		if err != nil {
			log.Cerror(err)
			time.Sleep(time.Second)
			continue
		}

		if thisInputs == lastInputs {
			continue
		}

		speedChan <- thisInputs
		lastInputs = thisInputs
	}
}

func (sv *Service) speedEval() {

	log.Printf("############### speedSensor eval started...")
	defer log.Printf("############### speedSensor eval stopped.")

	for speedSensor != nil {

		select {

		case <-time.After(time.Millisecond):

			now := time.Now()

			for pin := 0; pin < 16; pin++ {

				state := speedStates[pin]

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

				speedStates[pin] = state
			}

		case inputs, ok := <-speedChan:

			if !ok {
				return
			}

			now := time.Now()

			for pin := 0; pin < 16; pin++ {

				mask := uint16(1 << pin)
				active := inputs&mask != 0

				state := speedStates[pin]

				if state.active != active {
					state.active = active
					state.dirty = true
					state.time = now
					speedStates[pin] = state
				}
			}
		}
	}
}

var lightSensor *mcp23017.MCP23017

func (sv *Service) lightTest(sensor *mcp23017.MCP23017) {

	if lightSensor != nil {
		return
	}

	lightSensor = sensor

	go sv.lightLoop()
}

func (sv *Service) lightLoop() {

	for {

		_ = lightSensor.SetPinDirections(0x0000)

		for loop := 1; loop < 5; loop++ {

			_ = lightSensor.WritePins(0xffff)
			//vals, _ := lightSensor.ReadPins()
			time.Sleep(time.Millisecond * 250)

			_ = lightSensor.WritePins(0x0000)
			//vals, _ = lightSensor.ReadPins()
			time.Sleep(time.Millisecond * 250)
		}

		//_ = lightSensor.WritePin(5, mcp23017.PinLogicHi)
		//_ = lightSensor.WritePin(9, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = lightSensor.WritePin(0, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(1, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(2, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(3, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(4, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = lightSensor.WritePin(5, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(6, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(7, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(8, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(9, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = lightSensor.WritePin(10, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(11, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(12, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(13, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = lightSensor.WritePin(14, mcp23017.PinLogicHi)

		time.Sleep(time.Second * 5)
	}
}
