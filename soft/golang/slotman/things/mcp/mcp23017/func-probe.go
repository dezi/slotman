package mcp23017

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string, desiredAddresses []byte) (things []*MCP23017, err error) {

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

	for _, devicePath := range devicePaths {

		for _, address := range desiredAddresses {

			deviceAddrPath := fmt.Sprintf("%s:%02x", devicePath, address)

			if simple.StringInArray(busyDevicePaths, deviceAddrPath) {
				continue
			}

			log.Printf("Probing MCP23017 deviceAddrPath=%s", deviceAddrPath)

			i2cDev := i2c.NewDevice(devicePath, address)
			tryErr := i2cDev.Open()
			if tryErr != nil {
				log.Cerror(tryErr)
				continue
			}

			isThing := false

			value, tryErr := i2cDev.ReadRegByte(byte(RegisterIoConA))
			if tryErr == nil {

				log.Printf("Identified MCP23017 devicePath=%s value=%04x", deviceAddrPath, value)

				isThing = true
			}

			_ = i2cDev.Close()

			if isThing {
				things = append(things, NewMCP23017(deviceAddrPath))
			}
		}
	}

	return
}
