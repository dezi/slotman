package ld6001a

import (
	"slotman/drivers/impl/uart"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busySerialPaths []string) (things []*LD6001a, err error) {

	devicePaths, err := uart.GetDevicePaths()
	if err != nil {
		log.Cerror(err)
		return
	}

	for _, devicePath := range devicePaths {

		if simple.StringInArray(busySerialPaths, devicePath) {
			continue
		}

		if devicePath == "/dev/ttyS0" {
			continue
		}

		//if runtime.GOOS == "darwin" && !strings.HasPrefix(devicePath, "/dev/tty.usbserial") {
		//
		//	//
		//	// Exclude bogus devices on OSX to speed up testing.
		//	//
		//
		//	continue
		//}

		for _, baudRate := range baudRates {

			log.Printf("Probing LD6001a devicePath=%s baudRate=%d", devicePath, baudRate)

			se := NewLD6001a(devicePath, baudRate)

			se.isProbe = true

			tryErr := se.Open()
			if tryErr != nil {
				log.Cerror(tryErr)
				continue
			}

			isValid := false

			for try := 1; try <= 3; try++ {

				var date, version, uid string
				date, version, uid, tryErr = se.GetVersion()
				if tryErr == nil {
					log.Printf("Identified LD6001a devicePath=%s baudRate=%d", devicePath, baudRate)
					log.Printf("Identified LD6001a date=%s version=%s uid=%s ", date, version, uid)
					isValid = true
					break
				}
			}

			_ = se.Close()
			se.isProbe = false

			if isValid {
				things = append(things, se)
			}
		}
	}

	return
}
