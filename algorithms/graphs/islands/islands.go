package islands

// Given a 2D grid grid where '1' represents land and '0' represents water, count and return the number of islands.

// An island is formed by connecting adjacent lands horizontally or vertically and is surrounded by water. You may assume water is surrounding the grid (i.e., all the edges are water).

// Examples
// ___________________________________________________________________________________
// Example 1
// // input:[
//     ["0","1","1","1","0"],
//     ["0","1","0","1","0"],
//     ["1","1","0","0","0"],
//     ["0","0","0","0","0"]
//   ]
// output: 1

// Example 2:
// input: [
//   ["1","1","0","0","1"],
//   ["1","1","0","0","1"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// output: 4
// ___________________________________________________________________________________

type Pair struct {
	Row    int
	Column int
}

var directions = []Pair{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, // up
}

func NumIslands(grid [][]byte) int {
	numRows, numCols := len(grid), len(grid[0])

	islands := 0

	for rowIndex := range numRows {
		for colIndex := range numCols {
			if grid[rowIndex][colIndex] == '1' && grid[rowIndex][colIndex] != '0' {
				islands++
				dfs(Pair{rowIndex, colIndex}, grid)
			}
		}
	}

	return islands
}

func dfs(node Pair, grid [][]byte) {
	grid[node.Row][node.Column] = '0'
	numRows, numCols := len(grid), len(grid[0])

	for _, direction := range directions {
		neighbour := Pair{Row: node.Row + direction.Row, Column: node.Column + direction.Column}
		if isInBounds(neighbour, numRows, numCols) && grid[neighbour.Row][neighbour.Column] == '1' && grid[neighbour.Row][neighbour.Column] != '0' {
			dfs(neighbour, grid)
		}
	}
}

func isInBounds(node Pair, numRows, numCols int) bool {
	if node.Column < 0 || node.Row < 0 {
		return false
	}

	return node.Row < numRows && node.Column < numCols
}
