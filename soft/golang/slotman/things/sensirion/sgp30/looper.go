package sgp30

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SGP30) readLoop() {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	for se.IsStarted {

		time.Sleep(time.Millisecond * 1000)

	}
}
