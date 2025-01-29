package controls

import (
	"slotman/things/mcp/mcp23017"
	"slotman/things/pololu/mxt550"
	"slotman/things/ti/ads1115"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.loadCalibrations()
	sv.checkSpeedControls()
	sv.checkSensors()
}

func (sv *Service) loadCalibrations() {

	if sv.speedControlCalibrations != nil {
		return
	}

	log.Printf("Initializing speed control calibrations...")

	sv.speedControlCalibrations = make([]*SpeedControlCalibration, maxTracks)

	for track := range sv.speedControlCalibrations {
		sv.speedControlCalibrations[track] = &SpeedControlCalibration{
			MinValue: 7000,
			MaxValue: 32765,
		}
	}
}

func (sv *Service) checkSpeedControls() {

	if sv.speedControlChannels != nil {
		return
	}

	log.Printf("Initializing speed controls...")

	sv.speedControlAttached = make([]bool, maxTracks)
	sv.speedControlChannels = make([]chan uint16, maxTracks)

	for track := range sv.speedControlChannels {
		sv.speedControlChannels[track] = make(chan uint16, 3)
		go sv.speedControlHandler(track)
	}
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

	if sv.ads1115s == nil {

		sv.ads1115s, err = ads1115.ProbeThings(nil, []byte{0x48, 0x49})
		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sv.ads1115s {

				switch sensor.GetThingAddress() {

				case 0x48:
					sv.ads1115Device1 = sensor

				case 0x49:
					sv.ads1115Device2 = sensor

				default:
					continue
				}

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					continue
				}

				sensor.SetResendMs(250)

				_ = sensor.SetGain(0, ads1115.Gain2)
				_ = sensor.SetGain(1, ads1115.Gain2)
				_ = sensor.SetGain(2, ads1115.Gain2)
				_ = sensor.SetGain(3, ads1115.Gain2)

				_ = sensor.SetRate(0, ads1115.Rate860Sps)
				_ = sensor.SetRate(1, ads1115.Rate860Sps)
				_ = sensor.SetRate(2, ads1115.Rate860Sps)
				_ = sensor.SetRate(3, ads1115.Rate860Sps)

				_ = sensor.SetCapMin(0, 7000)
				_ = sensor.SetCapMin(1, 7000)
				_ = sensor.SetCapMin(2, 7000)
				_ = sensor.SetCapMin(3, 7000)

				err = sensor.Start()
				if err != nil {
					continue
				}
			}
		}
	}
}
