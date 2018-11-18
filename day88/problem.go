package day88

import (
	"errors"
)

var errDivideByZero = errors.New("can't divide by zero")

// ErrDivideByZero is the error returned if you attempt to divide by zero.
func ErrDivideByZero() error {
	return errDivideByZero
}

// BruteForceDivision calculates the division by looping subtraction.
// Runs in O(Quotient) time and O(1) space.
// Returns an error if divisor is zero.
func BruteForceDivision(dividend, divisor int32) (int32, error) {
	if divisor == 0 {
		return 0, ErrDivideByZero()
	}
	negative := (dividend < 0 && divisor > 0) || (dividend >= 0 && divisor < 0)
	dividend = abs(dividend)
	divisor = abs(divisor)
	quotient := int32(0)
	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}
	if negative {
		return -quotient, nil
	}
	return quotient, nil
}

// BitShiftDivision calculates the division by bit-shifting.
// Runs in O(1) time because it loops over 63-bits and O(1) space.
// Returns an error if divisor is zero.
func BitShiftDivision(dividend, divisor int32) (int32, error) {
	if divisor == 0 {
		return 0, ErrDivideByZero()
	}
	negative := (dividend < 0 && divisor > 0) || (dividend >= 0 && divisor < 0)
	dividend = abs(dividend)
	divisor = abs(divisor)
	var quotient, tmp int64
	for i := uint(31); i < 32; i-- {
		if tmp+(int64(divisor)<<i) <= int64(dividend) {
			tmp += int64(divisor) << i
			quotient |= int64(1) << i
		}
	}
	if negative {
		return -int32(quotient), nil
	}
	return int32(quotient), nil
}

// abs returns the absolute value.
func abs(val int32) int32 {
	if val < 0 {
		return -val
	}
	return val
}
