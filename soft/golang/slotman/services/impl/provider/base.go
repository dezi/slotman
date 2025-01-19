package provider

import (
	"errors"
	"fmt"
	"time"
)

func SetProvider(provider BaseProvider) {

	providerMutex.Lock()
	providers[provider.GetName()] = provider
	providerMutex.Unlock()

	controlMutex.Lock()

	if control, ok := provider.(ControlProvider); ok {

		interval := control.GetControlOptions()

		ct := &controlTask{}
		ct.nextDue = time.Now().UnixNano()
		ct.interval = interval

		controlTasks[provider.GetName()] = ct
	}

	controlMutex.Unlock()
}

func UnsetProvider(provider BaseProvider) {

	providerMutex.Lock()
	delete(providers, provider.GetName())
	providerMutex.Unlock()

	controlMutex.Lock()
	delete(controlTasks, provider.GetName())
	controlMutex.Unlock()
}

func GetProvider(providerName Provider) (provider BaseProvider, err error) {

	providerMutex.Lock()
	provider, ok := providers[providerName]
	providerMutex.Unlock()

	if !ok {
		err = ErrNotFound(providerName)
		return
	}

	return
}

func ErrNotFound(providerName Provider) (err error) {
	err = errors.New(fmt.Sprintf("provider <%s> not found", providerName))
	return
}
