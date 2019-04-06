package day202

// IsIntegerPalindrome returns true if the integer is a
// palindrome.
// Runs in O(digits) time and O(digits) space.
func IsIntegerPalindrome(i int) bool {
	var digits []int
	for i != 0 {
		digits = append(digits, i%10)
		i /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

// IsIntegerPalindromeFaster returns true if the integer is a
// palindrome.
// Runs in O(digits) time and O(1) space.
func IsIntegerPalindromeFaster(i int) bool {
	divisor := 1
	for temp := i; temp/divisor > 10; divisor *= 10 {
	}
	for i >= 10 {
		highest := i / divisor
		lowest := i % 10
		if highest != lowest {
			return false
		}
		i %= divisor
		i /= 10
		divisor /= 100
	}
	return true
}
