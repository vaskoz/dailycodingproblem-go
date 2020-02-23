package day458

import "strings"

// Direction represents the eight cardinal and intercardinal directions.
type Direction string

const (
	// N is North
	N Direction = "N"
	// NE is North-East
	NE Direction = "NE"
	// E is East
	E Direction = "E"
	// SE is South-East
	SE Direction = "SE"
	// S is South
	S Direction = "S"
	// SW is South-West
	SW Direction = "SW"
	// W is West
	W Direction = "W"
	// NW is North-West
	NW Direction = "NW"
)

// Position represents a point in Euclidean geometry.
type Position struct {
	X, Y int
}

// directions defines the relative coordinates of each Direction.
// nolint
var directions = map[Direction]Position{
	N:  {0, -1},
	NE: {1, -1},
	E:  {1, 0},
	SE: {1, 1},
	S:  {0, 1},
	SW: {-1, 1},
	W:  {-1, 0},
	NW: {-1, -1},
}

// IsValidRules answers if the rules are consistent.
func IsValidRules(rules []string) bool {
	positions := make(map[string]Position)

	for _, rule := range rules {
		parts := strings.Split(rule, " ")
		to, dir, from := parts[0], Direction(parts[1]), parts[2]
		toPos, toFound := positions[to]
		fromPos, fromFound := positions[from]

		switch delta := directions[dir]; {
		case toFound && fromFound:
			if !checkValid(fromPos, toPos, dir) {
				return false
			}
		case toFound && !fromFound:
			positions[from] = Position{toPos.X - delta.X, toPos.Y - delta.Y}
		case !toFound && fromFound:
			positions[to] = Position{fromPos.X + delta.X, fromPos.Y + delta.Y}
		default:
			positions[from] = Position{0, 0}
			positions[to] = directions[dir]
		}
	}

	return true
}

func checkValid(from, to Position, dir Direction) bool {
	return isNorth(from, to, dir) ||
		isNorthEast(from, to, dir) ||
		isEast(from, to, dir) ||
		isSouthEast(from, to, dir) ||
		isSouth(from, to, dir) ||
		isSouthWest(from, to, dir) ||
		isWest(from, to, dir) ||
		isNorthWest(from, to, dir)
}

func isNorth(from, to Position, dir Direction) bool {
	return dir == N && to.Y < from.Y
}

func isNorthEast(from, to Position, dir Direction) bool {
	return dir == NE && to.Y < from.Y && to.X > from.X
}

func isEast(from, to Position, dir Direction) bool {
	return dir == E && to.X > from.X
}

func isSouthEast(from, to Position, dir Direction) bool {
	return dir == SE && to.X > from.X && to.Y > from.Y
}

func isSouth(from, to Position, dir Direction) bool {
	return dir == S && to.Y > from.Y
}

func isSouthWest(from, to Position, dir Direction) bool {
	return dir == SW && to.X < from.X && to.Y > from.Y
}

func isWest(from, to Position, dir Direction) bool {
	return dir == W && to.X < from.X
}

func isNorthWest(from, to Position, dir Direction) bool {
	return dir == NW && to.Y < from.Y && to.X < from.X
}
