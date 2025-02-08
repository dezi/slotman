package speedo

import (
	"slotman/things/pololu/mxt550"
	"time"
)

func (sv *Service) motoronSafetyLoop(motoron *mxt550.MXT550) {

	var loops int64

	for !sv.doExit {

		time.Sleep(time.Millisecond * 1000)

		var voltageMv uint32
		var tryErr error

		for {

			time.Sleep(time.Millisecond * 10)

			voltageMv, tryErr = motoron.GetVinVoltageMv(5000, mxt550.Motoron550)
			if tryErr != nil {
				continue
			}

			if voltageMv < 40000 {
				break
			}
		}

		address := motoron.GetThingAddress()

		var tracks []int

		if address == 0x18 {
			tracks = []int{0, 1}
		}

		if address == 0x19 {
			tracks = []int{2, 3}
		}

		if address == 0x1a {
			tracks = []int{4, 5}
		}

		if address == 0x1b {
			tracks = []int{6, 7}
		}

		if tracks == nil {
			continue
		}

		if loops%10 == 0 {
			sv.rce.OnMotoronVoltage(tracks, voltageMv)
			//log.Printf("Motoron address=%02x tracks=%v voltageMv=%d", address, tracks, voltageMv)
		}

		loops++
	}
}
