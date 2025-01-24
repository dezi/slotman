package provider

import (
	"os"
	"slotman/utils/log"
	"sync"
)

var (
	initialized  bool
	doExit       bool
	controlGroup sync.WaitGroup

	providers     map[Service]BaseService
	providerMutex sync.Mutex

	controlTasks map[Service]*controlTask
	controlMutex sync.Mutex
)

func StartService() (err error) {

	if initialized {
		return
	}

	providers = make(map[Service]BaseService)
	controlTasks = make(map[Service]*controlTask)

	controlGroup.Add(1)
	go controlLoop()

	initialized = true

	hostName, _ := os.Hostname()

	log.Printf("Started service on host=%s.", hostName)

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
