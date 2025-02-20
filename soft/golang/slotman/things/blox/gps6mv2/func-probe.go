package gps6mv2

import (
	"runtime"
	"slotman/drivers/impl/uart"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
	"time"
)

func ProbeThings(busySerialPaths []string) (things []*GPS6MV2, err error) {

	devicePaths, err := uart.GetDevicePaths()
	if err != nil {
		log.Cerror(err)
		return
	}

	devicePaths = []string{"/dev/tty.usbserial-10"}

	for _, devicePath := range devicePaths {

		if simple.StringInArray(busySerialPaths, devicePath) {
			continue
		}

		if devicePath == "/dev/ttyS0" {
			continue
		}

		if runtime.GOOS == "darwin" && !strings.HasPrefix(devicePath, "/dev/tty.usbserial") {

			//
			// Exclude bogus devices on OSX to speed up testing.
			//

			continue
		}

		for _, baudRate := range baudRates {

			log.Printf("Probing GPS6MV2 devicePath=%s baudRate=%d", devicePath, baudRate)

			se := NewGPS6MV2(devicePath, baudRate)

			se.isProbe = true

			tryErr := se.Open()
			if tryErr != nil {
				log.Cerror(tryErr)
				continue
			}

			isValid := false

			select {
			case <-time.After(time.Second * 4):
			case line := <-se.results:
				isValid = strings.HasPrefix(line, "$")
			}

			_ = se.Close()
			se.isProbe = false

			if isValid {
				log.Printf("Identified GPS6MV2 devicePath=%s baudRate=%d", devicePath, baudRate)
				things = append(things, se)
			}
		}
	}

	return
}
