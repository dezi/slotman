package provider

import (
	"slotman/utils/log"
	"sync"
)

var (
	initialized  bool
	doExit       bool
	controlGroup sync.WaitGroup

	providers     map[Provider]BaseProvider
	providerMutex sync.Mutex

	controlTasks map[Provider]*controlTask
	controlMutex sync.Mutex
)

func StartService() (err error) {

	if initialized {
		return
	}

	providers = make(map[Provider]BaseProvider)
	controlTasks = make(map[Provider]*controlTask)

	controlGroup.Add(1)
	go controlLoop()

	initialized = true

	log.Printf("Started service.")

	return
}

func StopService() (err error) {

	if !initialized {
		return
	}

	log.Printf("Stopping service...")

	doExit = true
	controlGroup.Wait()

	log.Printf("Finished all subprocesses.")

	providers = nil
	controlTasks = nil
	initialized = false
	doExit = false

	log.Printf("Stopped service.")

	return
}
