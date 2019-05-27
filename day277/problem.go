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
			if !validBytes(10, 1, bits) {
				return false
			}
			bits = bits[16:]
		case bits[0:4] == "1110":
			if !validBytes(18, 2, bits) {
				return false
			}
			bits = bits[24:]
		case bits[0:5] == "11110":
			if !validBytes(26, 3, bits) {
				return false
			}
			bits = bits[32:]
		default:
			return false
		}
	}
	return true
}

func validBytes(maxlen, bytes int, bits string) bool {
	if len(bits) < maxlen {
		return false
	}
	start := 8
	for i := 0; i < bytes; i++ {
		if bits[start:start+2] != "10" {
			return false
		}
		start += 8
	}
	return true
}
