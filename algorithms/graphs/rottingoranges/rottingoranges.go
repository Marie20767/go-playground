package rottingoranges

// You are given a 2-D matrix grid. Each cell can have one of three possible values:

// 0 representing an empty cell
// 1 representing a fresh fruit
// 2 representing a rotten fruit

// Every minute, if a fresh fruit is horizontally or vertically adjacent to a rotten fruit, then the fresh fruit also becomes rotten.
// Return the minimum number of minutes that must elapse until there are zero fresh fruits remaining.
// If this state is impossible within the grid, return -1.

// Examples (see example picture rottingoranges.png):
// 1. input: [[1,1,0], [0,1,1],[0,1,2]] output: 4
// 2. input: [[1,0,1],[0,2,0],[1,0,1]] output: -1
// 3. input: [[0]] output: 0
// 4. input: [[1]] output: -1
// 5. input: [[2]] output: 1

type Node struct {
	Row    int
	Column int
}

var directions = []Node{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, // up
}

func isInBounds(node Node, numRows, numCols int) bool {
	if node.Column < 0 || node.Row < 0 {
		return false
	}

	return node.Row < numRows && node.Column < numCols
}

func OrangesRotting(grid [][]int) int {
	queue := []Node{}
	minutes := 0
	numRows, numCols := len(grid), len(grid[0])
	fresh := 0

	for rowIndex := range numRows {
		for colIndex := range numCols {
			if grid[rowIndex][colIndex] == 1 {
				fresh++
			}

			if grid[rowIndex][colIndex] == 2 {
				queue = append(queue, Node{Row: rowIndex, Column: colIndex})
			}
		}
	}

	for fresh > 0 && len(queue) > 0 {
		for range queue {
			first := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				neighbour := Node{Row: first.Row + direction.Row, Column: first.Column + direction.Column}

				if isInBounds(neighbour, numRows, numCols) && grid[neighbour.Row][neighbour.Column] == 1 {
					fresh--
					grid[neighbour.Row][neighbour.Column] = 2
					queue = append(queue, neighbour)
				}
			}
		}

		minutes++
	}

	if fresh == 0 {
		return minutes
	}

	return -1
}
