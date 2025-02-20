package sgp40

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SGP40) readLoop() {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	for se.IsStarted {

		time.Sleep(time.Millisecond * 1000)

		percent, err := se.MeasureAirQuality()
		if err != nil {
			log.Cerror(err)
			continue
		}

		handler := se.handler
		if handler != nil {
			go handler.OnAirQuality(se, percent)
		}
	}
}
