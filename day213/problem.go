package day213

import (
	"fmt"
	"strconv"
)

// ValidIPAddresses returns all the possible valid
// IPv4 addresses using all the digits in the string.
func ValidIPAddresses(digits string) []string {
	return validIPAddresses(digits, 0)
}

func validIPAddresses(digits string, octet int) []string {
	if octet == 4 {
		if len(digits) == 0 {
			return []string{""}
		}
		return nil
	}
	var one, two, three, result []string
	if len(digits) > 0 {
		one = validIPAddresses(digits[1:], octet+1)
	}
	if len(digits) > 1 && digits[0] != byte('0') {
		two = validIPAddresses(digits[2:], octet+1)
	}
	if len(digits) > 2 && digits[0] != byte('0') {
		if third, _ := strconv.Atoi(digits[:3]); third < 256 && third > 99 {
			three = validIPAddresses(digits[3:], octet+1)
		}
	}
	for i, ips := range [][]string{one, two, three} {
		for _, ip := range ips {
			if len(ip) != 0 {
				result = append(result, fmt.Sprintf("%s.%s", digits[:i+1], ip))
			} else {
				result = append(result, digits[:i+1])
			}
		}
	}
	return result
}
