package ambient

import (
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (sv *Service) OnThingOpened(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing opened uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnThingClosed(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing closed uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnThingStarted(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing started uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnThingStopped(thing things.Thing) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()
	log.Printf("Thing stopped uuid=%s vendor=<%s> model=<%s>", uuid[:8], vendor, short)
}

func (sv *Service) OnAirQuality(thing things.Thing, percent float64) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()

	sv.currentAirPercents += percent
	sv.currentAirSamples++

	if time.Now().Unix()-sv.lastAirTime.Unix() < 60 {
		return
	}

	sv.lastAirTime = time.Now()

	percent = sv.currentAirPercents / float64(sv.currentAirSamples)

	sv.currentAirPercents = 0
	sv.currentAirSamples = 0

	log.Printf("Thing air-quality uuid=%s vendor=<%s> model=<%s> percent=%0.1f", uuid[:8],
		vendor, short, percent)
}

func (sv *Service) OnTemperature(thing things.Thing, celsius float64) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()

	log.Printf("Thing temperature uuid=%s vendor=<%s> model=<%s> celsius=%0.1f", uuid[:8], vendor, short, celsius)
}

func (sv *Service) OnPressure(thing things.Thing, hPa float64) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()

	log.Printf("Thing pressure uuid=%s vendor=<%s> model=<%s> hPa=%0.1f", uuid[:8], vendor, short, hPa)
}

func (sv *Service) OnHumidity(thing things.Thing, percent float64) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()

	log.Printf("Thing humidity uuid=%s vendor=<%s> model=<%s> percent=%0.1f", uuid[:8], vendor, short, percent)
}

func (sv *Service) OnRGBColor(thing things.Thing, r, g, b, lux int) {

	uuid := thing.GetUuid()
	vendor, _, short := thing.GetModelInfo()

	log.Printf("Thing rgb-color uuid=%s vendor=<%s> model=<%s> r=%d g=%d b=%d lux=%d",
		uuid[:8], vendor, short, r, g, b, lux)
}
