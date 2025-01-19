package provider

import "time"

type Provider string

type controlTask struct {
	isInGo   bool
	nextDue  int64
	interval time.Duration
}

type BaseProvider interface {
	GetName() (name Provider)
}

type ControlProvider interface {
	GetName() (name Provider)
	GetControlOptions() (interval time.Duration)
	DoControlTask()
}
