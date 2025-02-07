package speedo

import (
	"slotman/things/pololu/mxt550"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkSensors()
}

func (sv *Service) checkSensors() {

	var err error

	if sv.mxt550s == nil {

		sv.mxt550s, err = mxt550.ProbeThings(nil, []byte{0x18, 0x19, 0x1a})

		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sv.mxt550s {

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					continue
				}

				sensor.SetDebug(false)

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

				switch sensor.GetThingAddress() {

				case 0x18:
					sv.mxt550Motoron1 = sensor

				case 0x19:
					sv.mxt550Motoron2 = sensor

				case 0x1a:
					sv.mxt550Motoron3 = sensor

				case 0x1b:
					sv.mxt550Motoron4 = sensor

				default:
					continue
				}
			}
		}
	}

}
