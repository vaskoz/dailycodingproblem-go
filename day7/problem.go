package day7

// NumberOfDecodings calculates the number of possible decodings where
// letters are encoding as a=1, b=2 ... z=26
func NumberOfDecodings(str string) int {
	if length := len(str); length == 0 || length == 1 {
		return 1
	}
	var count int
	if last := str[len(str)-1]; last > '0' && last <= '9' {
		count = NumberOfDecodings(str[:len(str)-1])
	}
	if second := str[len(str)-2]; second == '1' || second == '2' && str[len(str)-1] < '7' {
		count += NumberOfDecodings(str[:len(str)-2])
	}
	return count
}
