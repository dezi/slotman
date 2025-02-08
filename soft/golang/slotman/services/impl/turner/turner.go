package turner

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"slotman/services/iface/pilots"
	"slotman/services/iface/proxy"
	"slotman/services/iface/teams"
	"slotman/services/iface/turner"
	"slotman/services/impl/provider"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

type Service struct {
	prx proxy.Interface
	tms teams.Interface
	plt pilots.Interface

	turnDisplay1 *gc9a01.GC9A01
	turnDisplay2 *gc9a01.GC9A01

	fontRegular *truetype.Font

	faceRegularNormal font.Face
	faceRegularLarge  font.Face

	teamIndex  int
	pilotIndex int

	isProxyServer bool
	isProxyClient bool

	loopCount int
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.isProxyServer = simple.GetExecName() == "proxy"
	singleTon.isProxyClient = simple.GOOS == "darwin"

	singleTon.prx, err = proxy.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	if singleTon.isProxyClient {

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
	}

	singleTon.prx.Subscribe(AreaTurner, singleTon)

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	singleTon.prx.Unsubscribe(AreaTurner)

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return turner.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 5
	return
}
