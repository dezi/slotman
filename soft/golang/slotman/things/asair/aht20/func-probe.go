package aht20

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string) (things []*AHT20, err error) {

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

		log.Printf("Probing AHT20 deviceAddrPath=%s", deviceAddrPath)

		i2cDev := i2c.NewDevice(devicePath, ThingI2CAddress)
		tryErr := i2cDev.Open()
		if tryErr != nil {
			continue
		}

		isValid := false

		tryErr = i2cDev.WriteRegBytes(byte(RegisterInit), []byte{0x08, 0x00})
		if tryErr == nil {
			log.Printf("Identified AHT20 devicePath=%s", devicePath)
			isValid = true
		}

		_ = i2cDev.Close()

		if isValid {
			things = append(things, NewAHT20(deviceAddrPath))
		}
	}

	return
}
