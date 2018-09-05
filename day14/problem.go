package day14

import (
	"math/rand"
	"time"
)

// EstimatePi uses monte carlo simulation to estimate PI.
func EstimatePi(iterations int) float64 {
	rand.Seed(time.Now().UnixNano())
	var inCircle, total int
	for total = 0; total < iterations; total++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			inCircle++
		}
	}
	return 4.0 * float64(inCircle) / float64(total)
}
