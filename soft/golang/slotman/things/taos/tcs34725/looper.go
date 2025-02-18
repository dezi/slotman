package tcs34725

import (
	"math"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *TCS34725) readLoop() {

	if !se.IsStarted {
		err := things.ErrThingNotStarted
		log.Cerror(err)
		return
	}

	var lastEvent int64
	var lastR, lastG, lastB int
	var lastLux int

	for se.IsStarted {

		time.Sleep(time.Millisecond * 1000)

		r, g, b, lux, err := se.ReadRgbColor()
		if err != nil {
			log.Cerror(err)
			continue
		}

		if math.Abs(float64(lastR-r)) >= se.threshold ||
			math.Abs(float64(lastG-g)) >= se.threshold ||
			math.Abs(float64(lastB-b)) >= se.threshold ||
			math.Abs(float64(lastLux-lux)) >= se.threshold ||
			time.Now().Unix()-lastEvent >= 60 {

			if se.handler != nil {
				se.handler.OnRGBColor(se, r, g, b, lux)
			}

			lastR, lastG, lastB = r, g, b
			lastLux = lux
			lastEvent = time.Now().Unix()
		}
	}
}
