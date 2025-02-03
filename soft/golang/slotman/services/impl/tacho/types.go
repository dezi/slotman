package tacho

import "time"

type SpeedState struct {
	active bool
	dirty  bool
	time   time.Time
}

type SpeedRead struct {
	pinStates uint16
	readTime  time.Time
}

type TrackState struct {
	Round       int
	RoundMillis int
	RoundTs     time.Time

	SpeedKmh float64
	SpeedTS  time.Time

	IsAtStart bool
}
