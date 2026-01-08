package surrounded

// You are given a 2-D matrix board containing 'X' and 'O' characters.
// If a continous, four-directionally connected group of 'O's is surrounded by 'X's,
// it is considered to be surrounded.
// Change all surrounded regions of 'O's to 'X's and do so in-place by
// modifying the input board.

// Note that regions that are on the border are not considered surrounded regions.

// Examples (see example picture surrounded.png):

// Example 1:
// Input: board = [
//   ["X","X","X","X"],
//   ["X","X","O","X"],
//   ["X","X","X","X"],
//   ["X","X","X","O"]
// ]

// Output: board = [
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","X","X","O"]
// ]

// Example 2:
// Input: board = [
//   ["X","X","X","X"],
//   ["X","O","O","X"],
//   ["X","O","O","X"],
//   ["X","X","X","O"]
// ]

// Output: [
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","X","X","O"]
// ]

// Example 3:
// Input: board = [
//   ["O","O","X","X"],
//   ["X","O","O","X"],
//   ["X","O","O","X"],
//   ["X","X","X","O"]
// ]

// Output: board = [
//   ["O","X","X","X"],
//   ["X","O","O","X"],
//   ["X","O","O","X"],
//   ["X","X","X","O"]
// ]

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

func SurroundedRegions(board [][]byte) {
	numRows, numCols := len(board), len(board[0])

	// looping through top and bottom borders
	for i := range numCols {
		dfs(Pair{0, i}, board) // top
		dfs(Pair{numRows - 1, i}, board) // bottom
	}

	// looping through left and right borders
	for i := range numRows {
		dfs(Pair{i, 0}, board) // left
		dfs(Pair{i, numCols - 1}, board) // right
	}

	for rowIndex := range numRows {
		for colIndex := range numCols {
			switch board[rowIndex][colIndex] {
			case 'R':
				board[rowIndex][colIndex] = 'O'
			case 'O':
				board[rowIndex][colIndex] = 'X'
			default: // continue loop
			}
		}
	}
}

func dfs(node Pair, board [][]byte) {
	if board[node.Row][node.Column] != 'O' {
		return
	}

	board[node.Row][node.Column] = 'R'

	for _, direction := range directions {
		neighbour := Pair{node.Row + direction.Row, node.Column + direction.Column}

		if isInBounds(neighbour, board) {
			dfs(neighbour, board)
		}
	}
}

func isInBounds(node Pair, board [][]byte) bool {
	numRows, numCols := len(board), len(board[0])

	if node.Row < 0 || node.Column < 0 {
		return false
	}

	return node.Row < numRows && node.Column < numCols
}
