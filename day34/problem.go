package day34

// MakePalindrome calculates the minimum number of insertions to make a string a palindrome.
// Runs in O(N) time and O(N) extra space for favorLeft and favorRight.
func MakePalindrome(str string) string {
	letters := []rune(str)
	favorLeft := make([]rune, len(letters))
	copy(favorLeft, letters)
	favorRight := make([]rune, len(letters))
	copy(favorRight, letters)
	left, right := 0, len(favorLeft)-1
	for left < right {
		if favorLeft[left] != favorLeft[right] {
			// insert the left character on the right side
			favorLeft = append(favorLeft, 'x')
			copy(favorLeft[right+2:], favorLeft[right+1:])
			favorLeft[right+1] = favorLeft[left]
			right++
		}
		left++
		right--
	}
	left, right = 0, len(favorRight)-1
	for left < right {
		if favorRight[left] != favorRight[right] {
			// insert the right character on the left side
			favorRight = append(favorRight, 'x')
			copy(favorRight[left+1:], favorRight[left:])
			favorRight[left] = favorRight[right+1]
			right++
		}
		left++
		right--
	}
	return chooseResult(favorLeft, favorRight)
}

func chooseResult(left, right []rune) string {
	if len(left) < len(right) {
		return string(left)
	} else if len(left) == len(right) {
		if left[0] < right[0] {
			return string(left)
		}
		return string(right)
	} else {
		return string(right)
	}
}
