package aht20

import (
	"math"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *AHT20) readLoop() {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	var lastHumTime int64
	var lastCelTime int64
	var lastHumVal float64
	var lastCelVal float64

	for se.IsStarted {

		time.Sleep(time.Millisecond * 1000)

		humidity, celsius, err := se.ReadMeasurement()
		if err != nil {
			log.Cerror(err)
			continue
		}

		if math.Abs(lastHumVal-humidity) >= se.threshold || time.Now().Unix()-lastHumTime >= 60 {

			if se.handler != nil {
				se.handler.OnHumidity(se, humidity)
			}

			lastHumVal = humidity
			lastHumTime = time.Now().Unix()
		}

		if math.Abs(lastCelVal-celsius) >= se.threshold || time.Now().Unix()-lastCelTime >= 60 {

			if se.handler != nil {
				se.handler.OnTemperature(se, celsius)
			}

			lastCelVal = celsius
			lastCelTime = time.Now().Unix()
		}
	}
}
