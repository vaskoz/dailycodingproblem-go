package day376

// Pos is a position.
type Pos struct {
	Row, Col int
}

// ClosestCoinToMe returns the closest coin to my position
// using the Manhattan distance.
func ClosestCoinToMe(coins []Pos, me Pos) Pos {
	closest := Pos{}
	closestDistance := len(coins) * 2

	for _, coin := range coins {
		if dist := abs(me.Row-coin.Row) + abs(me.Col-coin.Col); dist < closestDistance {
			closestDistance = dist
			closest = coin
		}
	}

	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
