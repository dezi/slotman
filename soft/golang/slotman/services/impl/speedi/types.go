package speedi

type SpeedControlCalibration struct {
	MinValue uint16
	MaxValue uint16
}

type PlayerControlCurve struct {
	BrkPercent float64
	MinPercent float64
	MaxPercent float64
}
