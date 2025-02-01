package speedi

import (
	"slotman/services/iface/proxy"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/impl/provider"
	"slotman/things/ti/ads1115"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

type Service struct {
	prx proxy.Interface
	sdo speedo.Interface

	ads1115s []*ads1115.ADS1115

	ads1115Device1 *ads1115.ADS1115
	ads1115Device2 *ads1115.ADS1115

	speedControlAttached     []bool
	speedControlChannels     []chan uint16
	speedControlCalibrations []*SpeedControlCalibration

	isProxyServer bool
	isProxyClient bool

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

	singleTon.prx, err = proxy.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.sdo, err = speedo.GetInstance()
	if err != nil {
		log.Cerror(err)
		return
	}

	singleTon.isProxyServer = simple.GetExecName() == "proxy"
	singleTon.isProxyClient = simple.GOOS == "darwin"

	singleTon.prx.Subscribe(AreaSpeedi, singleTon)

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	singleTon.prx.Unsubscribe(AreaSpeedi)

	singleTon.doExit = true

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return speedi.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
