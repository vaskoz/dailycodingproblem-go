package day385

import (
	"encoding/hex"
	"strings"
)

// nolint
var freqMap = map[byte]float64{
	'a': 0.08167, 'b': 0.01492, 'c': 0.02782,
	'd': 0.04253, 'e': 0.12702, 'f': 0.02228,
	'g': 0.02015, 'h': 0.06094, 'i': 0.06966,
	'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
	'm': 0.02406, 'n': 0.06749, 'o': 0.07507,
	'p': 0.01929, 'q': 0.00095, 'r': 0.05987,
	's': 0.06327, 't': 0.09056, 'u': 0.02758,
	'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
	'y': 0.01974, 'z': 0.00074,
}

// DecryptSecretXorMessage checks all possible chars used for XOR.
func DecryptSecretXorMessage(msg string) string {
	decoded, err := hex.DecodeString(msg)
	if err != nil {
		panic("bad input message not hex")
	}

	smallest := 1000.0
	ans := ""

	for b := byte(0); b < byte(255); b++ {
		freq := make(map[byte]int)
		sb := strings.Builder{}
		totalC := 0
		invalidC := 0

		for i := 0; i < len(decoded); i++ {
			newB := decoded[i] ^ b
			lowered := strings.ToLower(string(newB))

			if lowerC := lowered[0]; lowerC >= 'a' && lowerC <= 'z' {
				freq[lowerC]++
			} else if lowerC != ' ' {
				invalidC++
			}
			totalC++

			sb.WriteByte(newB)
		}

		if invalidC > 0 {
			continue
		}

		var diff float64

		for k, v := range freq {
			diff += abs(freqMap[k] - float64(v)/float64(totalC))
		}

		if diff < smallest {
			smallest = diff
			ans = sb.String()
		}
	}

	return ans
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}

	return a
}
