package teams

import (
	"image"
	"slotman/services/iface/teams"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sync"
	"time"
)

type Service struct {
	teams map[simple.UUIDHex]*slotman.Team

	teamLogoFull  map[simple.UUIDHex]*image.RGBA
	teamLogoSmall map[simple.UUIDHex]*image.RGBA
	teamCarFull   map[simple.UUIDHex]*image.RGBA
	teamCarSmall  map[simple.UUIDHex]*image.RGBA

	mapsLock sync.Mutex
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.teams = make(map[simple.UUIDHex]*slotman.Team)

	singleTon.teamLogoFull = make(map[simple.UUIDHex]*image.RGBA)
	singleTon.teamLogoSmall = make(map[simple.UUIDHex]*image.RGBA)

	singleTon.teamCarFull = make(map[simple.UUIDHex]*image.RGBA)
	singleTon.teamCarSmall = make(map[simple.UUIDHex]*image.RGBA)

	singleTon.loadMockups()

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return teams.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
