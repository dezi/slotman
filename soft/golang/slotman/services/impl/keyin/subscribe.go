package keyin

import (
	"slotman/services/type/keyin"
)

func (sv *Service) Subscribe(subscriber keyin.Subscriber) {

	sv.subscribersLock.Lock()
	defer sv.subscribersLock.Unlock()

	sv.subscribers[subscriber] = true
}

func (sv *Service) Unsubscribe(subscriber keyin.Subscriber) {

	sv.subscribersLock.Lock()
	defer sv.subscribersLock.Unlock()

	delete(sv.subscribers, subscriber)
}
