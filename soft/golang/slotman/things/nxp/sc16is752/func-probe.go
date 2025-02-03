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

			//value, tryErr := i2cDev.ReadRegU16BE(byte(RegisterConfig))
			//if tryErr == nil {
			//
			//	log.Printf("Identified ADS1115 devicePath=%s value=%04x", deviceAddrPath, value)
			//
			//	isValid = true
			//}

			_ = i2cDev.Close()

			if isValid {
				things = append(things, NewSC15IS752(deviceAddrPath))
			}
		}
	}

	return
}
