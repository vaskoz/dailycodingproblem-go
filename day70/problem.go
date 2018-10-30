package day70

// PerfectNth returns the n-th perfect number.
// A perfect number is one where the digits sum to 10.
// A panic will happen if the argument is less than 1.
func PerfectNth(n int) int {
	if n < 1 {
		panic("n must be greater than 0")
	}
	perfectNum := 10 // 10 is the seed value.
	for n != 0 {
		perfectNum += 9
		digits := perfectNum
		sum := 0
		for digits != 0 {
			sum += (digits % 10)
			digits /= 10
		}
		if sum == 10 {
			n--
		}
	}
	return perfectNum
}
