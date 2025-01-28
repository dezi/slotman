package ads1115

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *ADS1115) readLoop() {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	var value uint16
	var tryErr error

	var lastTimes [4]int64
	var lastValues [4]uint16

	eps := uint16(16)
	abs := func(x, y uint16) uint16 {
		if x > y {
			return x - y
		} else {
			return y - x
		}
	}

	for se.IsStarted {

		time.Sleep(time.Millisecond)

		if se.handler == nil {
			continue
		}

		for input := 0; input <= 3; input++ {

			value, tryErr = se.ReadADConversion(input)
			if tryErr != nil {
				continue
			}

			if abs(lastValues[input], value) <= eps && time.Now().UnixMilli()-lastTimes[input] <= se.resendMs {
				continue
			}

			se.handler.OnADConversion(se, input, value)
			lastValues[input] = value
			lastTimes[input] = time.Now().UnixMilli()
		}
	}
}
