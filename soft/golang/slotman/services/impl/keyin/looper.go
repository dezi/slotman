package keyin

import (
	"slotman/utils/log"
)

func (sv *Service) looper() {

	log.Printf("Keyboard input reader started...")
	defer log.Printf("Keyboard input reader done.")

	input := make([]byte, 1)

	for !sv.doExit {

		consoleReader := sv.consoleReader
		if consoleReader == nil {
			return
		}

		xfer, tryErr := consoleReader.Read(input)
		if tryErr != nil {
			return
		}

		if xfer == 0 {
			continue
		}

		sv.subscribersLock.Lock()

		for subscriber := range sv.subscribers {
			subscriber.OnAsciiKeyPress(input[0])
		}

		sv.subscribersLock.Unlock()

		log.Printf("Keyboard input=%d", input[0])
	}
}
