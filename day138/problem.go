package day138

import "errors"

var errNotPossible = errors.New("Can not make this value")

// ErrNotPossible represents that making this currency with these
// denominations is impossible.
func ErrNotPossible() error {
	return errNotPossible
}

// MinCurrencyRequiredBrute returns the minimum number of currency
// necessary to make the desired amount.
func MinCurrencyRequiredBrute(amt int, denomCents []int) (int, error) {
	if amt < 0 {
		return 0, errNotPossible
	} else if amt == 0 {
		return 0, nil
	}
	min := int(^uint(0) >> 1)
	found := false
	for _, denom := range denomCents {
		if count, err := MinCurrencyRequiredBrute(amt-denom, denomCents); err == nil && count+1 < min {
			found = true
			min = count + 1
		}
	}
	if found {
		return min, nil
	}
	return 0, errNotPossible
}
