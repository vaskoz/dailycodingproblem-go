package day158

// Maze is a matrix of walls and free space.
// 1 - is a wall. 0 - is free space.
type Maze [][]int

// NumRoutesThruMaze returns the number of routes
// from the upper-left to the bottom-right.
func NumRoutesThruMaze(maze Maze) int {
	return numRoutesThruMaze(maze, 0, 0)
}

func numRoutesThruMaze(maze Maze, row, col int) int {
	if row >= len(maze) || col >= len(maze[0]) || maze[row][col] == 1 {
		return 0
	} else if row == len(maze)-1 && col == len(maze[row])-1 {
		return 1
	}
	return numRoutesThruMaze(maze, row+1, col) +
		numRoutesThruMaze(maze, row, col+1)
}
