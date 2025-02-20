package sgp40

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SGP40) readLoop() {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	//var lastHPaTime int64
	//var lastCelTime int64
	//var lastHPaVal float64
	//var lastCelVal float64

	for se.IsStarted {

		time.Sleep(time.Millisecond * 1000)

		//celsius, err := se.ReadTemperature()
		//if err != nil {
		//	log.Cerror(err)
		//} else {
		//
		//	if math.Abs(lastCelVal-celsius) >= se.threshold || time.Now().Unix()-lastCelTime >= 60 {
		//
		//		if se.handler != nil {
		//			se.handler.OnTemperature(se, celsius)
		//		}
		//
		//		lastCelVal = celsius
		//		lastCelTime = time.Now().Unix()
		//	}
		//}
		//
		//hPa, err := se.ReadPressure()
		//if err != nil {
		//	log.Cerror(err)
		//} else {
		//
		//	if math.Abs(lastHPaVal-hPa) >= se.threshold || time.Now().Unix()-lastHPaTime >= 60 {
		//
		//		if se.handler != nil {
		//			se.handler.OnPressure(se, hPa)
		//		}
		//
		//		lastHPaVal = hPa
		//		lastHPaTime = time.Now().Unix()
		//	}
		//}
	}
}
