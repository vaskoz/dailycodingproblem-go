package day472

// NumberOfDecodings calculates the number of possible decodings where
// letters are encoding as a=1, b=2 ... z=26
func NumberOfDecodings(str string) int {
	if length := len(str); length < 2 {
		return 1
	}

	count := 0

	if first := str[0]; first == '1' || (first == '2' && str[1] < '7') {
		count = NumberOfDecodings(str[2:])
	}

	return count + NumberOfDecodings(str[1:])
}
