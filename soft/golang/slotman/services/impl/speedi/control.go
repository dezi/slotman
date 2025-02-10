package speedi

import (
	"slotman/services/type/slotman"
	"slotman/things/ti/ads1115"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.loadCalibrations()
	sv.checkSpeedControls()
	sv.checkSensors()
}

func (sv *Service) checkSpeedControls() {

	if sv.speedControlChannels != nil {
		return
	}

	log.Printf("Initializing speed controls...")

	sv.speedControlAttached = make([]bool, slotman.MaxTracks)
	sv.speedControlChannels = make([]chan uint16, slotman.MaxTracks)

	for track := range sv.speedControlChannels {
		sv.speedControlChannels[track] = make(chan uint16, 3)
		go sv.speedControlHandler(track)
	}
}

func (sv *Service) loadCalibrations() {

	if sv.isProxyServer {
		return
	}

	if sv.speedControlCalibrations != nil {
		return
	}

	log.Printf("Initializing speed control calibrations...")

	sv.speedControlCalibrations = make([]*SpeedControlCalibration, slotman.MaxTracks)

	for track := range sv.speedControlCalibrations {
		sv.speedControlCalibrations[track] = &SpeedControlCalibration{
			MinValue: 7000,
			MaxValue: 32765,
		}
	}
}

func (sv *Service) checkSensors() {

	if sv.isProxyClient {
		return
	}

	var err error

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
