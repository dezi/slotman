package sc16is752

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SC15IS752) readLoop(channel byte) {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	for se.IsStarted {

		if se.pollSleep[channel] == 0 {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		time.Sleep(time.Millisecond * time.Duration(se.pollSleep[channel]))

		se.readLock.Lock()

		//
		// todo read data...
		//

		se.readLock.Unlock()
	}
}
