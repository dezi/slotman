package ld6001a

import (
	"slotman/drivers/impl/uart"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
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

		if devicePath != "/dev/i2c-1:48-0" {
			continue
		}

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

			_ = se.StopWorking()

			for try := 0; try < 3; try++ {
				tryErr = se.ReadParams()
				isValid = tryErr == nil && se.Params.SoftwareVersion != ""
				if isValid {
					break
				}
				time.Sleep(time.Second * 5)
			}

			_ = se.Close()
			se.isProbe = false

			if isValid {
				log.Printf("Identified LD6001A devicePath=%s baudRate=%d", devicePath, baudRate)
				log.Printf("Identified LD6001A version=%s", se.Params.SoftwareVersion)
				things = append(things, se)
			}
		}
	}

	return
}
