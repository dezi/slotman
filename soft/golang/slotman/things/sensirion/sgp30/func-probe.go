package sgp30

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string) (things []*SGP30, err error) {

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

		log.Printf("Probing SGP30 deviceAddrPath=%s", deviceAddrPath)

		i2cDev := i2c.NewDevice(devicePath, ThingI2CAddress)
		tryErr := i2cDev.Open()
		if tryErr != nil {
			log.Cerror(tryErr)
			continue
		}

		isValid := false

		xfer, tryErr := i2cDev.WriteBytes([]byte{0x36, 0x82})
		log.Cerror(tryErr)
		if tryErr == nil && xfer == 2 {
			log.Printf("Identified SGP30 devicePath=%s", devicePath)
			isValid = true
		}

		_ = i2cDev.Close()

		if isValid {
			things = append(things, NewBMP280(deviceAddrPath))
		}
	}

	return
}
