package controls

import (
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkSensors()
}

func (sv *Service) checkSensors() {

	var err error

	if sv.mcp23017s == nil {

		sv.mcp23017s, err = mcp23017.ProbeThings(nil, []byte{0x20, 0x21})

		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sv.mcp23017s {

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					log.Cerror(err)
					continue
				}

				err = sensor.Start()
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					continue
				}

				sv.mcp23017s = append(sv.mcp23017s, sensor)

				if sensor.GetThingAddress() == 0x20 {
					sv.mcp23017StartLight = sensor
					sv.lightTest(sensor)
				}

				if sensor.GetThingAddress() == 0x21 {
					sv.mcp23017SpeedMeasure = sensor
					sv.speedTest(sensor)
				}

				log.Printf("Registered sensor MCP23017 path=%s uuid=%s",
					sensor.DevicePath, sensor.GetUuid()[:8])
			}
		}
	}
}
