package day303

import "math"

// AngleClockHands returns the nearest angle between clock hands.
// Works at minute granularity.
func AngleClockHands(hh, mm int) int {
	if hh < 0 || hh >= 24 || mm < 0 || mm >= 60 {
		panic("invalid time")
	}
	if hh >= 12 {
		hh -= 12
	}
	hourAngle := 360.0 * float64(hh) / 12.0
	minuteAngle := 360.0 * float64(mm) / 60.0
	diffAngle := math.Abs(hourAngle - minuteAngle)
	if diffAngle > 180 {
		diffAngle -= 360
	}
	return int(math.Round(math.Abs(diffAngle)))
}
