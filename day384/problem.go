package day384

import (
	"errors"
	"sort"
)

// nolint
var errImpossibleTarget = errors.New("impossible to make target change with denominations")

const MaxInt = int(^uint(0) >> 1)

// FewestCoinsBrute returns the fewest required coins to make target
// in change.
// Returns an error if it's not possible.
func FewestCoinsBrute(denom []int, target int) (int, error) {
	copied := append([]int{}, denom...)
	sort.Sort(sort.Reverse(sort.IntSlice(copied)))

	return fewestCoinsBrute(copied, target, MaxInt)
}

func fewestCoinsBrute(denom []int, target, coinsSoFar int) (int, error) {
	if target < 0 || coinsSoFar < 0 {
		return 0, errImpossibleTarget
	} else if target == 0 {
		return 0, nil
	}

	smallest := MaxInt

	for _, d := range denom {
		if res, err := fewestCoinsBrute(denom, target-d, coinsSoFar-1); err == nil && res < smallest {
			smallest = res
			coinsSoFar = res
		}
	}

	if smallest == MaxInt {
		return 0, errImpossibleTarget
	}

	return 1 + smallest, nil
}
