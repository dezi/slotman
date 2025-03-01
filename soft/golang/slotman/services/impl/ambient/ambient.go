package ambient

import (
	"slotman/services/iface/ambient"
	"slotman/services/impl/provider"
	"slotman/things/asair/aht20"
	"slotman/things/bosch/bmp280"
	"slotman/things/sensirion/sgp30"
	"slotman/things/sensirion/sgp40"
	"slotman/things/taos/tcs34725"
	"slotman/utils/log"
	"time"
)

type Service struct {
	sgp30    *sgp30.SGP30
	sgp40    *sgp40.SGP40
	aht20    *aht20.AHT20
	bmp280   *bmp280.BMP280
	tcs34725 *tcs34725.TCS34725

	currentAirPercents float64
	currentAirSamples  int

	lastAirTime time.Time

	doExit bool
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	singleTon.doExit = true

	if singleTon.sgp40 != nil {
		_ = singleTon.sgp40.Close()
		singleTon.sgp40 = nil
	}

	if singleTon.aht20 != nil {
		_ = singleTon.aht20.Close()
		singleTon.aht20 = nil
	}

	if singleTon.bmp280 != nil {
		_ = singleTon.bmp280.Close()
		singleTon.bmp280 = nil
	}

	if singleTon.tcs34725 != nil {
		_ = singleTon.tcs34725.Close()
		singleTon.tcs34725 = nil
	}

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return ambient.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
