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

		xfer, tryErr := i2cDev.WriteBytes([]byte{0x36, 0x82})
		//xfer, tryErr := i2cDev.WriteBytes([]byte{0x28, 0x0E})
		log.Cerror(tryErr)
		if tryErr == nil && xfer == 2 {
			log.Printf("Identified SGP40 devicePath=%s", devicePath)

			time.Sleep(time.Millisecond * 500)

			result := make([]byte, 9)
			xfer, tryErr = i2cDev.ReadBytes(result)

			log.Printf("############ ReadBytes xfer=%d result=[ %02x ] tryErr=%v",
				xfer, result, tryErr)

			log.Printf("############# crc=%02x", calculateCrc(result[0:2]))
			log.Printf("############# crc=%02x", calculateCrc(result[3:5]))
			log.Printf("############# crc=%02x", calculateCrc(result[6:8]))
			isValid = true
		}

		_ = i2cDev.Close()

		if isValid {
			things = append(things, NewBMP280(deviceAddrPath))
		}
	}

	return
}
