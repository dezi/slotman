package proxy

import "slotman/utils/log"

func (sv *Service) DoControlTask() {
	sv.checkServer()
}

func (sv *Service) checkServer() {
	err := sv.startServers()
	log.Cerror(err)
}
