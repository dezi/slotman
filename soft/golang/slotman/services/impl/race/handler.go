package race

import (
	"slotman/utils/log"
)

func (sv *Service) OnAmpelClickShort() {
	log.Printf("OnAmpelClickShort...")
}

func (sv *Service) OnAmpelClickLong() {
	log.Printf("OnAmpelClickLong...")
}

func (sv *Service) OnMotoronVoltage(tracks []int, voltageMv uint32) {
	log.Printf("OnMotoronVoltage tracks=%v voltageMv=%d", tracks, voltageMv)
}
