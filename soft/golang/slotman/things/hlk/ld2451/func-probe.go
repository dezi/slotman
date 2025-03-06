package ld2451

import (
	"runtime"
	"slotman/drivers/impl/uart"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func ProbeThings(busySerialPaths []string) (things []*LD2451, err error) {

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

		if runtime.GOOS == "darwin" && !strings.HasPrefix(devicePath, "/dev/tty.usbserial") {

			//
			// Exclude bogus devices on OSX to speed up testing.
			//

			continue
		}

		for _, baudRate := range baudRates {

			log.Printf("Probing LD2451 devicePath=%s baudRate=%d", devicePath, baudRate)

			se := NewLD2451(devicePath, baudRate)

			se.isProbe = true

			tryErr := se.Open()
			if tryErr != nil {
				log.Cerror(tryErr)
				continue
			}

			isValid := false

			for try := 1; try <= 3; try++ {
				var protocol int
				protocol, tryErr = se.EnableConfigurations()
				if tryErr == nil {

					_ = se.FactoryReset()

					//_ = se.RestartModule()
					//_, _ = se.EnableConfigurations()

					_ = se.SetDetectionParams(
						10, 0, 0,
						DetectionModeBoth)

					maxDist, minSpeed, delay, mode, _ := se.GetDetectionParams()
					trigger, noise, _ := se.GetSensitivityParams()

					fType, major, minor, _ := se.GetVersion()

					_ = se.EndConfiguration()

					log.Printf("Identified LD2451 devicePath=%s baudRate=%d", devicePath, baudRate)
					log.Printf("Identified LD2451 protocol=%04x", protocol)

					log.Printf("Identified LD2451 fType=%04x major=%04x minor=%04x",
						fType, major, minor)

					log.Printf("Identified LD2451 maxDist=%d minSpeed=%d delay=%d mode=%d",
						maxDist, minSpeed, delay, mode)

					log.Printf("Identified LD2451 trigger=%d noise=%d",
						trigger, noise)

					isValid = true
					break
				}
			}

			_ = se.Close()
			se.isProbe = false
			_ = se.Close()
			se.isProbe = false

			if isValid {
				things = append(things, se)
				break
			}
		}
	}

	return
}
