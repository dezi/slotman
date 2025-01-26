package bmp280

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string) (things []*BMP280, err error) {

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

		log.Printf("Probing BMP280 deviceAddrPath=%s", deviceAddrPath)

		i2cDev := i2c.NewDevice(devicePath, ThingI2CAddress)
		tryErr := i2cDev.Open()
		if tryErr != nil {
			continue
		}

		isValid := false

		id, tryErr := i2cDev.ReadRegByte(0xd0)
		if tryErr == nil && id == 0x58 {
			log.Printf("Identified BMP280 devicePath=%s id=%02x", devicePath, id)
			isValid = true
		}

		_ = i2cDev.Close()

		if isValid {
			things = append(things, NewBMP280(deviceAddrPath))
		}
	}

	return
}
