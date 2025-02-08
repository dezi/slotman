package race

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/speedo"
	"time"
)

func (sv *Service) DoControlTask() {
	sv.checkServices()
	sv.checkLooper()
}

func (sv *Service) checkServices() {

	if sv.servicesReady {
		return
	}

	//
	// Yield time for target services to register.
	//

	time.Sleep(time.Second)

	var tryErr error

	sv.amp, tryErr = ampel.GetInstance()
	if tryErr != nil {
		return
	}

	sv.sdo, tryErr = speedo.GetInstance()
	if tryErr != nil {
		return
	}

	sv.servicesReady = true
}

func (sv *Service) checkLooper() {

	if !sv.servicesReady || sv.looperStarted {
		return
	}

	sv.looperStarted = true

	go sv.looper()
}
