package day23

import (
	"errors"
	"sort"
)

// Board with false representing open space and true representing a wall.
type Board [][]bool

// Cell represents a row & col position on the board.
type Cell struct {
	row, col int
}

var errNoPath = errors.New("No path possible")

// MinimumSteps performs depth-first-search to find a path.
func MinimumSteps(board Board, start, end Cell, visited map[Cell]struct{}) (int, error) {
	visited[start] = struct{}{}
	if start == end {
		delete(visited, start)
		return 0, nil
	}
	var results []int
	cells := []Cell{{start.row + 1, start.col}, {start.row - 1, start.col}, {start.row, start.col - 1}, {start.row, start.col + 1}}
	for _, cell := range cells {
		if _, found := visited[cell]; !found && isValidNextCell(board, cell) {
			dist, err := MinimumSteps(board, cell, end, visited)
			if err == nil {
				results = append(results, dist)
			}
		}
	}
	sort.Ints(results)
	delete(visited, start)
	if len(results) == 0 {
		return 0, errNoPath
	}
	return 1 + results[0], nil
}

// isValidNextCell checks if the cell is a valid cell to move to.
// Valid means inbounds and also not a wall (true)
func isValidNextCell(board Board, cell Cell) bool {
	return cell.row >= 0 && cell.row < len(board) &&
		cell.col >= 0 && cell.col < len(board[cell.row]) &&
		!board[cell.row][cell.col]
}
