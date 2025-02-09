package turner

const (
	goldColor  = "e0bf78"
	circleSize = 70
)

var (
	oddPositions = []float64{
		120 - circleSize/2, 240 - circleSize,
		circleSize / 3.6, 240 - circleSize/2.0 - circleSize,
		240 - circleSize/3.6 - circleSize, 240 - circleSize/2.0 - circleSize,
		circleSize / 12.5, 240 - circleSize*2.52,
		240 - circleSize/12.5 - circleSize, 240 - circleSize*2.52,
		circleSize / 1.45, 240 - circleSize*3.33,
		240 - circleSize/1.45 - circleSize, 240 - circleSize*3.33,
	}

	evenPositions = []float64{
		circleSize / 1.45, 240 - circleSize/5.5 - circleSize,
		240 - circleSize/1.45 - circleSize, 240 - circleSize/5.5 - circleSize,
		circleSize / 12.5, 240 - circleSize/0.9 - circleSize,
		240 - circleSize/12.5 - circleSize, 240 - circleSize/0.9 - circleSize,
		circleSize / 2, 240 - circleSize*3.1,
		240 - circleSize/2 - circleSize, 240 - circleSize*3.1,
	}
)
