package day148

import "fmt"

// GrayCodes returns a slice of strings representing the bits of
// gray codes for the requested number of bits.
func GrayCodes(bits int) []string {
	if bits < 1 {
		return nil
	} else if bits == 1 {
		return []string{"0", "1"}
	}
	smaller := GrayCodes(bits - 1)
	result := make([]string, 0, 2*len(smaller))
	for _, entry := range smaller {
		result = append(result, fmt.Sprintf("0%s", entry))
	}
	for i := range smaller {
		result = append(result, fmt.Sprintf("1%s", smaller[len(smaller)-1-i]))
	}
	return result
}
