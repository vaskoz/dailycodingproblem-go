package day127

// RevLLNum represents a reversed linked list of digits.
type RevLLNum struct {
	Digit int
	Next  *RevLLNum
}

// SumRevLLNum returns the sum of two RevLLNum as a new RevLLNum.
func SumRevLLNum(first, second *RevLLNum) *RevLLNum {
	var carryOne bool
	result := &RevLLNum{}
	current := result
	for first != nil || second != nil {
		current.Next = &RevLLNum{}
		current = current.Next
		var sum int
		if first == nil {
			sum = second.Digit
			second = second.Next
		} else if second == nil {
			sum = first.Digit
			first = first.Next
		} else {
			sum = first.Digit + second.Digit
			first = first.Next
			second = second.Next
		}
		if carryOne {
			sum++
		}
		carryOne = sum > 9
		current.Digit = sum % 10
	}
	if carryOne {
		current.Next = &RevLLNum{1, nil}
	}
	return result.Next
}
