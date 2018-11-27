package day96

// HeapsAlgorithm returns a channel of integer slices.
func HeapsAlgorithm(digits []int) <-chan []int {
	num := 1
	for i := 2; i <= len(digits); i++ {
		num *= i
	}
	copied := make([]int, len(digits))
	copy(copied, digits)
	perms := make(chan []int, num)
	go generate(len(copied), copied, perms)
	return perms
}

func generate(n int, digits []int, perms chan []int) {
	if n == 1 {
		copied := make([]int, len(digits))
		copy(copied, digits)
		perms <- copied
		return
	}
	for i := 0; i < n-1; i++ {
		generate(n-1, digits, perms)
		if n%2 == 0 {
			digits[i], digits[n-1] = digits[n-1], digits[i]
		} else {
			digits[0], digits[n-1] = digits[n-1], digits[0]
		}
	}
	generate(n-1, digits, perms)
	if n == len(digits) {
		close(perms)
	}
}
