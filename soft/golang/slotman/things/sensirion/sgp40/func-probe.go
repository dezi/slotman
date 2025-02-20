package sgp40

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

func ProbeThings(busyDevicePaths []string) (things []*SGP40, err error) {

	//
	// Probe on i2c interfaces.
	//

	devicePaths, err := i2c.GetDevicePaths()
	if err != nil {
		log.Cerror(err)
		return
	}

	for _, devicePath := range devicePaths {

		deviceAddrPath := fmt.Sprintf("%s:%02x", devicePath, ThingI2CAddress)

		if simple.StringInArray(busyDevicePaths, deviceAddrPath) {
			continue
		}

		log.Printf("Probing SGP40 deviceAddrPath=%s", deviceAddrPath)

		i2cDev := i2c.NewDevice(devicePath, ThingI2CAddress)
		tryErr := i2cDev.Open()
		if tryErr != nil {
			log.Cerror(tryErr)
			continue
		}

		isValid := false

		//
		// Try self test command.
		//

		xfer, tryErr := i2cDev.WriteBytes([]byte{0x28, 0x0E})
		if tryErr == nil && xfer == 2 {

			//
			// Self test needs time to evaluate.
			//

			time.Sleep(time.Millisecond * 300)

			result := make([]byte, 3)
			xfer, tryErr = i2cDev.ReadBytes(result)

			if tryErr == nil && xfer == 3 {
				if result[2] == calculateCrc(result[0:2]) && result[0] == 0xd4 {
					log.Printf("Identified SGP40 devicePath=%s result=[ %02x ]", devicePath, result)
					isValid = true
				}
			}
		}

		_ = i2cDev.Close()

		if isValid {
			things = append(things, NewSGP40(deviceAddrPath))
		}
	}

	return
}
