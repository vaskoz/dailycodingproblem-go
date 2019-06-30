package day312

// TilingBoardWays returns the number of ways a board of size
// 2 x N can be tiled using dominoes and trominoes.
func TilingBoardWays(n int) int {
	switch {
	case n < 1:
		panic("n must be greater than 0")
	case n == 1:
		return 1
	case n == 2:
		return 2
	default:
		return TilingBoardWays(n-1) + TilingBoardWays(n-2)
	}
}
