package day252

import "math"

// EgyptianFractions convert a fraction into a list of denominators
// for an Egyptian fraction.
func EgyptianFractions(numerator, denominator int) []int {
	var denominators []int
	for numerator != 0 {
		d := int(math.Ceil(float64(denominator) / float64(numerator)))
		denominators = append(denominators, d)
		numerator = d*numerator - denominator
		denominator *= d
	}
	return denominators
}
