package day229

// Jump represents the movement that happens either up or down.
// Ladders give you a boost up, while a Snake lowers your position.
type Jump struct {
	From, To int
}

// SnakesAndLaddersFewestTurns by playing all games.
func SnakesAndLaddersFewestTurns(ladders, snakes []Jump, maxDepth int) int {
	jumps := make(map[int]int, len(ladders)+len(snakes))
	for _, j := range ladders {
		jumps[j.From] = j.To
	}
	for _, j := range snakes {
		jumps[j.From] = j.To
	}
	return snakesAndLaddersFewestTurns(jumps, 0, 0, int(^uint(0)>>1), maxDepth)
}

func snakesAndLaddersFewestTurns(jumps map[int]int, pos, depth, bestSoFar, maxDepth int) int {
	if pos == 100 {
		return depth
	} else if depth > maxDepth || depth > bestSoFar {
		return bestSoFar
	}
	for roll := 6; roll >= 0; roll-- {
		next := pos + roll
		if jumpPos, found := jumps[next]; found && jumpPos < pos { // avoid all snakes
			continue
		} else if found && jumpPos > pos {
			next = jumpPos
		}
		if next > 100 {
			break
		}
		if turns := snakesAndLaddersFewestTurns(jumps, next, depth+1, bestSoFar, maxDepth); turns < bestSoFar {
			bestSoFar = turns
		}
	}
	return bestSoFar
}
