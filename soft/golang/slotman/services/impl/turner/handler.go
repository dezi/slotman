package turner

import (
	"slotman/things"
	"slotman/utils/log"
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
