package proxy

import "slotman/services/type/proxy"

func (sv *Service) Subscribe(area proxy.Area, handler proxy.Subscriber) {

	sv.subscribersLock.Lock()
	defer sv.subscribersLock.Unlock()

	sv.subscribers[area] = handler
}

func (sv *Service) Unsubscribe(area proxy.Area) {

	sv.subscribersLock.Lock()
	defer sv.subscribersLock.Unlock()

	delete(sv.subscribers, area)

}
