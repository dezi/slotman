package mxt550

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string, desiredAddresses []byte) (things []*MXT550, err error) {

	//
	// Probe on i2c interfaces.
	//

	devicePaths, err := i2c.GetDevicePaths()
	if err != nil {
		log.Cerror(err)
		return
	}

	if desiredAddresses == nil {
		desiredAddresses = ThingAddresses
	}

	var usedAddresses []byte

	for _, devicePath := range devicePaths {

		for _, address := range desiredAddresses {

			deviceAddrPath := fmt.Sprintf("%s:%02x", devicePath, address)

			if simple.StringInArray(busyDevicePaths, deviceAddrPath) {
				continue
			}

			log.Printf("Probing MXT550 deviceAddrPath=%s", deviceAddrPath)

			mxt550 := NewMXT550(deviceAddrPath)

			tryErr := mxt550.Open()
			if tryErr != nil {
				continue
			}

			err = mxt550.DisableCrc()
			log.Cerror(err)

			productId, firmwareVersion, tryErr := mxt550.GetFirmwareVersion()
			if tryErr != nil {
				log.Cerror(tryErr)
				_ = mxt550.Close()
				continue
			}

			log.Printf("Identified MXT550 devicePath=%s productId=%04x firmwareVersion=%s",
				deviceAddrPath, productId, firmwareVersion)

			usedAddresses = append(usedAddresses, address)

			if address == 0x0f {

				log.Printf("Writable MXT550 devicePath=%s productId=%04x firmwareVersion=%s",
					deviceAddrPath, productId, firmwareVersion)

				var newAddress byte
				var oldAddress byte

				for _, customAddress := range CustomAddresses {

					present := false

					for _, usedAddress := range usedAddresses {
						if usedAddress == customAddress {
							present = true
							break
						}
					}

					if !present {
						newAddress = customAddress
						break
					}
				}

				if newAddress == 0 {

					log.Printf("All MXT550 custom addresses in use...")

				} else {

					oldAddress, err = mxt550.ReadEepromDeviceNumber()
					if err != nil {
						log.Cerror(err)
						return
					}

					log.Printf("Old MXT550 address is 0x%02x", oldAddress)
					log.Printf("New MXT550 custom address will be 0x%02x", newAddress)

					if oldAddress == newAddress {

						log.Printf("Already set up...")

					} else {

						err = mxt550.WriteEepromDeviceNumber(newAddress)
						if err != nil {
							log.Cerror(err)
							return
						}

						err = mxt550.Reset()
						if err != nil {
							log.Cerror(err)
							return
						}

						log.Printf("New MXT550 custom address set to 0x%02x", newAddress)
					}
				}

				_ = mxt550.Close()
				continue
			}

			_ = mxt550.Close()

			things = append(things, mxt550)
		}
	}

	return
}
