package day312

const panicMessage = "n must be greater than 0"

// TilingBoardWays returns the number of ways a board of size
// 2 x N can be tiled using dominoes and trominoes.
func TilingBoardWays(n int) int {
	switch {
	case n < 1:
		panic(panicMessage)
	case n == 1:
		return 1
	case n == 2:
		return 2
	default:
		return TilingBoardWays(n-1) + TilingBoardWays(n-2)
	}
}

// TilingBoardWaysFaster returns the number of ways a board of size
// 2 x N can be tiled using dominoes and trominoes.
func TilingBoardWaysFaster(n int) int {
	if n < 1 {
		panic("n must be greater than 0")
	}
	n1, n2 := 0, 1
	for i := 1; i <= n; i++ {
		tmp := n2
		n2 = n1 + n2
		n1 = tmp
	}
	return n2
}
