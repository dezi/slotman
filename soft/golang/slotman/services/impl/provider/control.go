package provider

import "time"

func controlLoop() {

	for !doExit {

		time.Sleep(time.Millisecond * 10)

		now := time.Now().UnixNano()

		var bestDue int64
		var bestTask *controlTask
		var bestProvider Provider

		controlMutex.Lock()

		for provider, ct := range controlTasks {

			if ct.nextDue >= now {
				continue
			}

			if bestTask == nil || ct.nextDue < bestDue {
				bestDue = ct.nextDue
				bestTask = ct
				bestProvider = provider
			}
		}

		controlMutex.Unlock()

		if bestTask != nil {

			bestTask.nextDue = now + int64(bestTask.interval)

			providerMutex.Lock()
			baseProvider := providers[bestProvider]
			providerMutex.Unlock()

			controlProvider, ok := baseProvider.(ControlProvider)

			if !ok {
				continue
			}

			if bestTask.isInGo {
				bestTask.nextDue = now + int64(time.Second)
				continue
			}

			go func() {
				bestTask.isInGo = true
				controlProvider.DoControlTask()
				bestTask.isInGo = false
			}()
		}

		if doExit {
			break
		}
	}

	controlGroup.Done()
}
