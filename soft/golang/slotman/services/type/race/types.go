package race

import (
	"slotman/services/type/slotman"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type RaceState string

//goland:noinspection GoNameStartsWithPackageName
const (
	RaceStateIdle          RaceState = "state.idle"
	RaceStateRaceWaiting   RaceState = "state.race.waiting"
	RaceStateRaceStarting  RaceState = "state.race.start"
	RaceStateRaceRunning   RaceState = "state.race.running"
	RaceStateRaceSuspended RaceState = "state.race.suspended"
	RaceStateRaceFinished  RaceState = "state.race.finished"
)

//goland:noinspection GoNameStartsWithPackageName
type RaceRecord struct {
	Pilot *slotman.Pilot

	Rounds   int
	Position int

	ActRound float64
	TopRound float64

	ActSpeed float64
	TopSpeed float64

	LastRoundTime time.Time
}
