package day352

import "reflect"

// Square represents either a white or black square
// on a crossword puzzle.
type Square int

const (
	// W is a white square
	W Square = iota
	// B is a white square
	B
)

// IsValidCrossword returns true if the puzzle
// meets all the requirements of a crossword puzzle.
// This is a brute force approach.
func IsValidCrossword(puzzle [][]Square) bool {
	if isValid := checkHorizontal(puzzle); !isValid {
		return false
	}
	if isValid := checkVertical(puzzle); !isValid {
		return false
	}
	if isValid := checkReachable(puzzle); !isValid {
		return false
	}
	if isValid := checkRotationallySymmetric(puzzle); !isValid {
		return false
	}
	return true
}

func checkRotationallySymmetric(puzzle [][]Square) bool {
	for r := range puzzle {
		rBottom := len(puzzle) - r - 1
		cRight := len(puzzle[r]) - r - 1
		cMost := len(puzzle[r]) - 1
		for i := 0; i < cRight; i++ {
			if puzzle[r][i] != puzzle[rBottom][cMost-i] {
				return false
			}
		}
	}
	return true
}

func checkReachable(puzzle [][]Square) bool {
	var startR, startC int
	whiteSquares := make(map[int]map[int]struct{})
	visited := make(map[int]map[int]struct{})
	for r := range puzzle {
		if whiteSquares[r] == nil {
			whiteSquares[r] = make(map[int]struct{})
			visited[r] = make(map[int]struct{})
		}
		for c := range puzzle[r] {
			if puzzle[r][c] == W {
				whiteSquares[r][c] = struct{}{}
				startR, startC = r, c
			}
		}
	}
	dfs(whiteSquares, visited, startR, startC)
	return reflect.DeepEqual(whiteSquares, visited)
}

func dfs(whiteSquares, visited map[int]map[int]struct{}, r, c int) {
	visited[r][c] = struct{}{}
	// up
	if _, exists := whiteSquares[r-1]; exists {
		if _, next := whiteSquares[r-1][c]; next {
			if _, seen := visited[r-1][c]; !seen {
				dfs(whiteSquares, visited, r-1, c)
			}
		}
	}
	// down
	if _, exists := whiteSquares[r+1]; exists {
		if _, next := whiteSquares[r+1][c]; next {
			if _, seen := visited[r+1][c]; !seen {
				dfs(whiteSquares, visited, r+1, c)
			}
		}
	}
	// left
	if _, exists := whiteSquares[r]; exists {
		if _, next := whiteSquares[r][c-1]; next {
			if _, seen := visited[r][c-1]; !seen {
				dfs(whiteSquares, visited, r, c-1)
			}
		}
	}
	// right
	if _, exists := whiteSquares[r]; exists {
		if _, next := whiteSquares[r][c+1]; next {
			if _, seen := visited[r][c+1]; !seen {
				dfs(whiteSquares, visited, r, c+1)
			}
		}
	}
}

func checkHorizontal(puzzle [][]Square) bool {
	var r, c int
	for r < len(puzzle) {
		if puzzle[r][c] == B {
			c++
		} else if puzzle[r][c] == W {
			count := 0
			for c < len(puzzle[r]) && puzzle[r][c] == W {
				count++
				c++
			}
			if count < 3 {
				return false
			}
		}
		if c == len(puzzle[r]) {
			c = 0
			r++
		}
	}
	return true
}

func checkVertical(puzzle [][]Square) bool {
	var r, c int
	for c < len(puzzle[r]) {
		if puzzle[r][c] == B {
			r++
		} else if puzzle[r][c] == W {
			count := 0
			for r < len(puzzle) && puzzle[r][c] == W {
				count++
				r++
			}
			if count < 3 {
				return false
			}
		}
		if r == len(puzzle) {
			r = 0
			c++
		}
	}
	return true
}
