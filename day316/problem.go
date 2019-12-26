package day316

import (
	"sort"
)

// CoinDenominationsNeeded returns the required denominations
// to achieve the number of ways to make change.
func CoinDenominationsNeeded(ways []int) []int {
	denoms := make(map[int]struct{})

	for d, way := range ways {
		if others := waysToProduce(d, denoms); (way == 1 && d > 0) || others == way-1 {
			denoms[d] = struct{}{}
		}
	}

	res := make([]int, 0, len(denoms))

	for k := range denoms {
		res = append(res, k)
	}

	sort.Ints(res)

	return res
}

func waysToProduce(target int, denoms map[int]struct{}) int {
	if _, found := denoms[target]; found || target == 0 {
		return 1
	}

	ways := 0

	for k := range denoms {
		if target%k == 0 {
			ways += waysToProduce(target/k, denoms)
		}
	}

	return ways
}
