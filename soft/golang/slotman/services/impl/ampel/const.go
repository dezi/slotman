package ampel

//goland:noinspection GoNameStartsWithPackageName
type AmpelState string

//goland:noinspection GoNameStartsWithPackageName
const (
	AmpelStateIdle        AmpelState = "ampel.idle"
	AmpelStateRaceStart   AmpelState = "ampel.race.start"
	AmpelStateRaceSuspend AmpelState = "ampel.race.suspend"
	AmpelStateRaceRestart AmpelState = "ampel.race.restart"
)

var (
	pinsRed    = []int{0, 1, 2, 3, 4}
	pinsGreen  = []int{5, 6, 7, 8, 9}
	pinsYellow = []int{10, 11, 12, 13, 14}
)
