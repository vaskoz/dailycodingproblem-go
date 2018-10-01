package day39

import (
	"math"
	"strings"
)

// Coord represents a coordinate in the game of life.
type Coord struct {
	X, Y int64
}

// GameOfLife represents the game state.
// '*' represents a live cell and '.' represents a dead cell.
type GameOfLife struct {
	living                 map[Coord]struct{}
	minX, minY, maxX, maxY int64
}

// NewGameOfLife parses a starting board.
// The given board's lower left coordinate with be 0,0
func NewGameOfLife(board string) *GameOfLife {
	gol := &GameOfLife{living: make(map[Coord]struct{})}
	rows := strings.Split(board, "\n")
	gol.maxY = int64(len(rows))
	for y := 0; y < len(rows); y++ {
		for x, val := range rows[len(rows)-1-y] {
			if val == '*' {
				gol.living[Coord{int64(x), int64(y)}] = struct{}{}
				if int64(x) >= gol.maxX {
					gol.maxX = int64(x) + 1
				}
			}
		}
	}
	gol.minX, gol.minY = -1, -1
	return gol
}

// String returns the string representation of the board now.
func (gol *GameOfLife) String() string {
	var sb strings.Builder
	for y := gol.maxY - 1; y > gol.minY; y-- {
		for x := gol.minX + 1; x < gol.maxX; x++ {
			if _, alive := gol.living[Coord{x, y}]; alive {
				sb.WriteRune('*')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// Step executes a single step of the simulation.
func (gol *GameOfLife) Step() {
	var birth, death []Coord
	for coord := range gol.living {
		if gol.willDie(coord) {
			death = append(death, coord)
		}
		birth = append(birth, gol.willBeBorn(coord)...)
	}
	for _, d := range death {
		delete(gol.living, d)
	}
	for _, b := range birth {
		gol.living[b] = struct{}{}
	}
	gol.minX, gol.minY, gol.maxX, gol.maxY = math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for coord := range gol.living {
		if coord.X <= gol.minX {
			gol.minX = coord.X - 1
		}
		if coord.X >= gol.maxX {
			gol.maxX = coord.X + 1
		}
		if coord.Y <= gol.minY {
			gol.minY = coord.Y - 1
		}
		if coord.Y >= gol.maxY {
			gol.maxY = coord.Y + 1
		}
	}
}

func (gol *GameOfLife) livingNeighbors(coord Coord) int {
	var livingNeighbors int
	for x := coord.X - 1; x <= coord.X+1; x++ {
		for y := coord.Y - 1; y <= coord.Y+1; y++ {
			if x == coord.X && y == coord.Y {
				continue
			}
			if _, alive := gol.living[Coord{x, y}]; alive {
				livingNeighbors++
			}
		}
	}
	return livingNeighbors
}

func (gol *GameOfLife) willBeBorn(coord Coord) []Coord {
	var birth []Coord
	checked := make(map[Coord]struct{})
	for x := coord.X - 1; x <= coord.X+1; x++ {
		for y := coord.Y - 1; y <= coord.Y+1; y++ {
			if x == coord.X && y == coord.Y {
				continue
			}
			coord := Coord{x, y}
			if _, alive := gol.living[coord]; alive {
				continue
			}
			if _, done := checked[coord]; !done {
				checked[coord] = struct{}{}
				living := gol.livingNeighbors(coord)
				if living == 3 {
					birth = append(birth, coord)
				}
			}
		}
	}
	return birth
}

func (gol *GameOfLife) willDie(coord Coord) bool {
	living := gol.livingNeighbors(coord)
	return living < 2 || living > 3
}
