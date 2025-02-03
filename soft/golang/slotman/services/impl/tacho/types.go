package tacho

import "time"

type SpeedState struct {
	active bool
	dirty  bool
	round  int
	time   time.Time
}
