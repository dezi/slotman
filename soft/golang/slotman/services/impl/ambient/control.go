package ambient

import (
	"slotman/things/asair/aht20"
	"slotman/things/bosch/bmp280"
	"slotman/things/sensirion/sgp30"
	"slotman/things/sensirion/sgp40"
	"slotman/things/taos/tcs34725"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	//sv.checkThingSGP30()
	sv.checkThingSGP40()
	sv.checkThingAHT20()
	sv.checkThingBMP280()
	sv.checkThingTCS34725()
}

func (sv *Service) checkThingSGP30() {

	if sv.sgp30 != nil {
		return
	}

	sensors, err := sgp30.ProbeThings(nil)

	if err != nil {
		log.Cerror(err)
		return
	}

	_ = sensors
}

func (sv *Service) checkThingSGP40() {

	if sv.sgp40 != nil {
		return
	}

	sensors, err := sgp40.ProbeThings(nil)

	if err != nil {
		log.Cerror(err)
		return
	}

	for _, sensor := range sensors {

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

		log.Printf("Registered air-quality SGP40 path=%s uuid=%s",
			sensor.DevicePath, sensor.GetUuid()[:8])

		sv.sgp40 = sensor
	}
}

func (sv *Service) checkThingAHT20() {

	if sv.aht20 != nil {
		return
	}

	sensors, err := aht20.ProbeThings(nil)

	if err != nil {
		log.Cerror(err)
		return
	}

	for _, sensor := range sensors {

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

		log.Printf("Registered humidity AHT20 path=%s uuid=%s",
			sensor.DevicePath, sensor.GetUuid()[:8])

		sv.aht20 = sensor
	}
}

func (sv *Service) checkThingBMP280() {

	if sv.bmp280 != nil {
		return
	}

	sensors, err := bmp280.ProbeThings(nil)

	if err != nil {
		log.Cerror(err)
		return
	}

	for _, sensor := range sensors {

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

		log.Printf("Registered pressure AHT20 path=%s uuid=%s",
			sensor.DevicePath, sensor.GetUuid()[:8])

		sv.bmp280 = sensor
	}
}

func (sv *Service) checkThingTCS34725() {

	if sv.tcs34725 != nil {
		return
	}

	sensors, err := tcs34725.ProbeThings(nil)

	if err != nil {
		log.Cerror(err)
		return
	}

	for _, sensor := range sensors {

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

		log.Printf("Registered rgb-color TCS34725 path=%s uuid=%s",
			sensor.DevicePath, sensor.GetUuid()[:8])

		sv.tcs34725 = sensor
	}
}
