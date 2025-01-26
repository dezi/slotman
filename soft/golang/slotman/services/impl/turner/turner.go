package turner

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"slotman/services/iface/pilots"
	"slotman/services/iface/teams"
	"slotman/services/iface/turner"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
	"time"
)

type Service struct {
	tms teams.Interface
	plt pilots.Interface

	turnDisplay1 *gc9a01.GC9A01
	turnDisplay2 *gc9a01.GC9A01

	fontRegular *truetype.Font

	faceRegularNormal font.Face
	faceRegularLarge  font.Face

	teams     []*slotman.Team
	teamIndex int

	pilots []*slotman.Pilot
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.tms, err = teams.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.plt, err = pilots.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.fontRegular, _ = truetype.Parse(goregular.TTF)

	singleTon.faceRegularNormal = truetype.NewFace(
		singleTon.fontRegular,
		&truetype.Options{Size: 24})

	singleTon.faceRegularLarge = truetype.NewFace(
		singleTon.fontRegular,
		&truetype.Options{Size: 40})

	singleTon.teams = singleTon.tms.GetAllTeams()

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
	return turner.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 10
	return
}
