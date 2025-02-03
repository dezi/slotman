package sc16is752

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string, desiredAddresses []byte) (things []*SC15IS752, err error) {

	//
	// Probe on i2c interfaces.
	//

	devicePaths, err := i2c.GetDevicePaths()
	if err != nil {
		log.Cerror(err)
		return
	}

	if desiredAddresses == nil {
		desiredAddresses = ThingI2CAddresses
	}

	for _, devicePath := range devicePaths {

		for _, address := range desiredAddresses {

			deviceAddrPath := fmt.Sprintf("%s:%02x", devicePath, address)

			if simple.StringInArray(busyDevicePaths, deviceAddrPath) {
				continue
			}

			//log.Printf("Probing ADS1115 deviceAddrPath=%s", deviceAddrPath)

			i2cDev := i2c.NewDevice(devicePath, address)
			tryErr := i2cDev.Open()
			if tryErr != nil {
				continue
			}

			isValid := false

			//
			// We probe using the user value register.
			//

			tryErr = i2cDev.WriteRegByte(RegSPR<<3, 0x55)
			if tryErr != nil {
				_ = i2cDev.Close()
				continue
			}

			var value byte

			value, tryErr = i2cDev.ReadRegByte(RegSPR << 3)
			if tryErr != nil {
				_ = i2cDev.Close()
				continue
			}

			//
			// Check if what we wrote is what we read.
			//

			isValid = value == 0x55

			_ = i2cDev.Close()

			if isValid {
				things = append(things, NewSC15IS752(deviceAddrPath))
			}
		}
	}

	return
}
