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

// TimesOverlappingHands returns all the [hh:mm] pairs with overlapping
// clock hands by checking all possible 12*60 hour/minute possibilities.
func TimesOverlappingHands() [][]int {
	var result [][]int
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if angle := AngleClockHands(h, m); angle == 0 {
				result = append(result, []int{h, m})
			}
		}
	}
	return result
}
