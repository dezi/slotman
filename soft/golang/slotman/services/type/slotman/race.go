package slotman

import (
	"slotman/utils/simple"
	"time"
)

type RaceState string

const (
	RaceStateIdle          RaceState = "state.idle"
	RaceStateRaceWaiting   RaceState = "state.race.waiting"
	RaceStateRaceStarting  RaceState = "state.race.start"
	RaceStateRaceRunning   RaceState = "state.race.running"
	RaceStateRaceSuspended RaceState = "state.race.suspended"
	RaceStateRaceFinished  RaceState = "state.race.finished"
)

type Race struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Title string `json:"title"`

	Tracks int `json:"tracks"`
	Rounds int `json:"rounds"`
}

type RaceInfo struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Track    int `json:"track"`
	Rounds   int `json:"rounds"`
	Position int `json:"position"`

	ActRound float64 `json:"actRound"`
	TopRound float64 `json:"topRound"`

	ActSpeed float64 `json:"actSpeed"`
	TopSpeed float64 `json:"topSpeed"`

	PilotUuid simple.UUIDHex `json:"pilotUuid"`

	LastRoundTime time.Time `json:"-"`
}

type TrackState int

const (
	TrackStateInactive TrackState = 0
	TrackStateActive   TrackState = 1
	TrackStateReady    TrackState = 2
	TrackStateFinished TrackState = 3
)
