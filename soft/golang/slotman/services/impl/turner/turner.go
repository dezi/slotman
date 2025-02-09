package turner

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
	"slotman/services/iface/pilots"
	"slotman/services/iface/proxy"
	"slotman/services/iface/race"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/iface/tacho"
	"slotman/services/iface/teams"
	"slotman/services/iface/turner"
	"slotman/services/impl/provider"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

type Service struct {
	rce race.Interface
	prx proxy.Interface
	tms teams.Interface
	plt pilots.Interface
	sdi speedi.Interface
	sdo speedo.Interface
	tco tacho.Interface

	turnDisplay1 *gc9a01.GC9A01
	turnDisplay2 *gc9a01.GC9A01

	fontBold    *truetype.Font
	fontRegular *truetype.Font

	faceBoldNormal font.Face
	faceBoldLarge  font.Face

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

	singleTon.prx.Subscribe(AreaTurner, singleTon)

	provider.SetProvider(singleTon)

	if singleTon.isProxyServer {
		return
	}

	singleTon.sdi, err = speedi.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.sdo, err = speedo.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.tco, err = tacho.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.rce, err = race.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

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

	singleTon.fontBold, _ = truetype.Parse(gobold.TTF)
	singleTon.fontRegular, _ = truetype.Parse(goregular.TTF)

	singleTon.faceBoldNormal = truetype.NewFace(
		singleTon.fontBold,
		&truetype.Options{Size: 26})

	singleTon.faceBoldLarge = truetype.NewFace(
		singleTon.fontBold,
		&truetype.Options{Size: 32})

	singleTon.faceRegularNormal = truetype.NewFace(
		singleTon.fontRegular,
		&truetype.Options{Size: 26})

	singleTon.faceRegularLarge = truetype.NewFace(
		singleTon.fontRegular,
		&truetype.Options{Size: 32})

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
