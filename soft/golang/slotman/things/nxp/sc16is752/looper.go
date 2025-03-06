package sc16is752

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SC15IS752) readLoop(channel byte) {

	defer se.loopGroup.Done()

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	var input []byte
	var startTime int64

	for se.IsStarted {

		if se.pollSleep[channel] == 0 {
			time.Sleep(time.Millisecond * 1)
		} else {
			time.Sleep(time.Millisecond * time.Duration(se.pollSleep[channel]))
		}

		data, err := se.ReadUartBytesNow(channel)
		if err != nil {
			continue
		}

		if len(input) == 0 {
			startTime = time.Now().UnixMilli()
		}

		input = append(input, data...)

		if len(input) > 0 && (len(input) > 128 || time.Now().UnixMilli()-startTime > 50) {

			//log.Printf("##### channel=%d len=%d input=[%02x]", channel, len(input), input)

			handler := se.handler
			if handler != nil {
				go handler.OnUartDataReceived(se, channel, input)
			}

			input = nil
		}
	}
}
