package tcs34725

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func ProbeThings(busyDevicePaths []string) (things []*TCS34725, err error) {

	//
	// Probe on i2c interfaces.
	//

	devicePaths, err := i2c.GetDevicePaths()
	if err != nil {
		log.Cerror(err)
		return
	}

	for _, devicePath := range devicePaths {

		deviceAddrPath := fmt.Sprintf("%s:%02x", devicePath, ThingAddress)

		if simple.StringInArray(busyDevicePaths, deviceAddrPath) {
			continue
		}

		log.Printf("Probing TCS34725 deviceAddrPath=%s", deviceAddrPath)

		i2cDev := i2c.NewDevice(devicePath, ThingAddress)
		tryErr := i2cDev.Open()
		if tryErr != nil {
			log.Cerror(tryErr)
			continue
		}

		isThing := false

		id, tryErr := i2cDev.ReadRegByte(byte(RegisterId))
		if tryErr == nil && (id == 0x4d || id == 0x44 || id == 0x10) {

			it, _ := i2cDev.ReadRegByte(byte(RegisterATime))
			gain, _ := i2cDev.ReadRegByte(byte(RegisterControl))

			log.Printf("Identified TCS34725 devicePath=%s id=%02x it=%02x gain=%02x",
				devicePath, id, it, gain)

			isThing = true
		}

		_ = i2cDev.Close()

		if isThing {
			things = append(things, NewTCS34275(deviceAddrPath))
		}
	}

	return
}
