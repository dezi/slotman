package speedo

import (
	"slotman/things/pololu/mxt550"
	"slotman/utils/log"
	"time"
)

func (sv *Service) motoronSafetyLoop(motoron *mxt550.MXT550) {

	var loops int64

	for !sv.doExit {

		time.Sleep(time.Millisecond * 1000)

		voltageMv, err := motoron.GetVinVoltageMv(5000, mxt550.Motoron550)
		if err != nil {
			log.Cerror(err)
			continue
		}

		address := motoron.GetThingAddress()

		tracks := ""

		if address == 0x18 {
			tracks = "1+2"
		}

		if address == 0x19 {
			tracks = "3+4"
		}

		if address == 0x1a {
			tracks = "5+6"
		}

		if tracks == "" {
			continue
		}

		if loops%10 == 0 {
			log.Printf("Motoron address=%02x tracks=%s voltageMv=%d", address, tracks, voltageMv)
		}

		loops++
	}
}
