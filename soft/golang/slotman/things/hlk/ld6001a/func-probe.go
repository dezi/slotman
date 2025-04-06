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

		if devicePath != "/dev/i2c-1:48-0+" && // 0bfe9a0a
			devicePath != "/dev/i2c-1:48-1+" && // 5b255252
			devicePath != "/dev/i2c-1:49-0" && // 4a3b77e2
			devicePath != "/dev/i2c-1:49-1" && // a325d0dd
			devicePath != "/dev/i2c-1:4c-0" && // ffcd0b16
			devicePath != "/dev/i2c-1:4c-1" && // 499f7a07
			devicePath != "/dev/i2c-1:4d-0+" && // 5ac050f5
			devicePath != "/dev/i2c-1:4d-1+" { // 1d6707a7
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

			tryErr = se.Reset()
			log.Cerror(tryErr)
			time.Sleep(time.Millisecond * 250)

			for try := 0; try < 5; try++ {
				tryErr = se.StopWorking()
				if tryErr == nil {
					break
				}
				log.Cerror(tryErr)
				time.Sleep(time.Millisecond * 100)
			}

			for try := 0; try < 5; try++ {
				tryErr = se.ReadParams()
				isValid = tryErr == nil && se.Params.SoftwareVersion != ""
				if isValid {
					break
				}
				time.Sleep(time.Second * 1)
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
