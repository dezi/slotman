package provider

import "time"

type Service string

type controlTask struct {
	isInGo   bool
	nextDue  int64
	interval time.Duration
}

type BaseService interface {
	GetName() (name Service)
}

type ControlService interface {
	GetName() (name Service)
	GetControlOptions() (interval time.Duration)
	DoControlTask()
}
