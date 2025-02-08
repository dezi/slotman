package ampel

import (
	"time"
)

func (sv *Service) buttonLoop() {

	var lastButtonDown bool
	var lastButtonTime int64

	for !sv.doExit {

		time.Sleep(time.Millisecond * 50)

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			break
		}

		sv.ampelLock.Lock()
		pins, err := ampelGpio.ReadPins()
		sv.ampelLock.Unlock()

		if err != nil {
			continue
		}

		thisButtonDown := pins&0x8000 == 0
		thisButtonTime := time.Now().UnixMilli()

		if lastButtonDown == thisButtonDown {
			continue
		}

		lastButtonDown = thisButtonDown

		if thisButtonDown {
			lastButtonTime = thisButtonTime
			sv.OnButtonPinDown(ampelGpio)
			continue
		}

		sv.OnButtonPinUp(ampelGpio)

		duration := thisButtonTime - lastButtonTime
		lastButtonTime = thisButtonTime

		if duration < 1000 {
			sv.OnButtonClickShort(ampelGpio)
		} else {
			sv.OnButtonClickLong(ampelGpio)
		}
	}
}
