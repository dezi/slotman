package proxy

import (
	"errors"
	"os"
	"slotman/services/type/proxy"
)

func (sv *Service) getTarget() (target string, err error) {

	hostName, err := os.Hostname()
	if err != nil {
		return
	}

	target, ok := proxy.ProxyTargets[hostName]
	if ok {
		return
	}

	err = errors.New("no proxy target for host")
	return
}
