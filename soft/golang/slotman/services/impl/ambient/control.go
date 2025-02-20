package ambient

import (
	"slotman/things/sensirion/sgp40"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkSensors()
}

func (sv *Service) checkSensors() {

	if sv.sgp40Co2 == nil {

		sensors, err := sgp40.ProbeThings(nil)

		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sensors {

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					log.Cerror(err)
					continue
				}

				_, _ = sensor.DoSelfTest()
				_, _ = sensor.ReadSerial()

				err = sensor.Start()
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					continue
				}

				_ = sensor.SetTemperature(2)

				log.Printf("Registered co2 sensor SGP40 path=%s uuid=%s",
					sensor.DevicePath, sensor.GetUuid()[:8])

				sv.sgp40Co2 = sensor
			}
		}
	}
}
