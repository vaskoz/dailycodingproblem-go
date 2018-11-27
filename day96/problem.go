package day96

// HeapsAlgorithmRecursive finds all the permutations recursively.
func HeapsAlgorithmRecursive(digits []int) <-chan []int {
	num := 1
	for i := 2; i <= len(digits); i++ {
		num *= i
	}
	copied := make([]int, len(digits))
	copy(copied, digits)
	perms := make(chan []int, num)
	go generateRecursive(len(copied), copied, perms)
	return perms
}

func generateRecursive(n int, digits []int, perms chan []int) {
	if n == 1 {
		copied := make([]int, len(digits))
		copy(copied, digits)
		perms <- copied
		return
	}
	for i := 0; i < n-1; i++ {
		generateRecursive(n-1, digits, perms)
		if n%2 == 0 {
			digits[i], digits[n-1] = digits[n-1], digits[i]
		} else {
			digits[0], digits[n-1] = digits[n-1], digits[0]
		}
	}
	generateRecursive(n-1, digits, perms)
	if n == len(digits) {
		close(perms)
	}
}

// HeapsAlgorithmIterative finds all the permutations iteratively.
func HeapsAlgorithmIterative(digits []int) <-chan []int {
	num := 1
	for i := 2; i <= len(digits); i++ {
		num *= i
	}
	copied := make([]int, len(digits))
	copy(copied, digits)
	perms := make(chan []int, num)
	go generateIterative(len(copied), copied, perms)
	return perms
}

func generateIterative(n int, digits []int, perms chan []int) {
	c := make([]int, len(digits))
	send := make([]int, len(digits))
	copy(send, digits)
	perms <- send
	var i int
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				digits[0], digits[i] = digits[i], digits[0]
			} else {
				digits[c[i]], digits[i] = digits[i], digits[c[i]]
			}
			send = make([]int, len(digits))
			copy(send, digits)
			perms <- send
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	close(perms)
}
