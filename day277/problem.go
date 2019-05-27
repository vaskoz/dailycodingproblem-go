package day277

// IsValidUTF8 returns true if the bits represent valid 1-4 byte
// UTF8.
// NOTE: Using a string of 0's and 1's to represent bits.
func IsValidUTF8(bits string) bool {
	if len(bits)%8 != 0 {
		panic("bytes should be multiples of 8-bits")
	}
	for bits != "" {
		switch {
		case bits[0] == '0':
			bits = bits[8:]
		case bits[0:3] == "110":
			if len(bits) < 10 || bits[8:10] != "10" {
				return false
			}
			bits = bits[16:]
		case bits[0:4] == "1110":
			if len(bits) < 18 || bits[8:10] != "10" || bits[16:18] != "10" {
				return false
			}
			bits = bits[24:]
		case bits[0:5] == "11110":
			if len(bits) < 26 || bits[8:10] != "10" || bits[16:18] != "10" || bits[24:26] != "10" {
				return false
			}
			bits = bits[32:]
		default:
			return false
		}
	}
	return true
}
