package controls

import (
	"slotman/things/mcp/mcp23017"
	"slotman/things/pololu/mxt550"
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

	if sv.mxt550s == nil {

		sv.mxt550s, err = mxt550.ProbeThings(nil, []byte{0x18, 0x19, 0x1a})

		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sv.mxt550s {

				switch sensor.GetThingAddress() {

				case 0x18:
					sv.mxt550Motoron1 = sensor

				case 0x19:
					sv.mxt550Motoron2 = sensor

				case 0x1a:
					sv.mxt550Motoron3 = sensor

				default:
					continue
				}

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					continue
				}

				err = sensor.Start()
				if err != nil {
					continue
				}

				err = sensor.ClearResetFlag()
				log.Cerror(err)

				err = sensor.ClearMotorFaultUnconditional()
				log.Cerror(err)

				//err = sensor.SetMaxAcceleration(1, 100)
				//log.Cerror(err)
				//err = sensor.SetMaxAcceleration(2, 100)
				//log.Cerror(err)
				//
				//err = sensor.SetMaxDeceleration(1, 200)
				//log.Cerror(err)
				//err = sensor.SetMaxDeceleration(2, 200)
				//log.Cerror(err)

				go sv.motoronSafetyLoop(sensor)
			}
		}
	}
}
