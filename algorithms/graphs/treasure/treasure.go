package treasure

// You are given a 2D grid initialized with these three possible values:

// -1 - A water cell that can not be traversed.
// 0 - A treasure chest.
// INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.
// Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest then the value should remain INF.

// Assume the grid can only be traversed up, down, left, or right.
// Modify the grid in-place.

// Examples:
// Example 1: Input: [
//   [0,-1],
//   [L,L]
// ]

// Output: [
//   [0,-1],
//   [1,2]
// ]

// Example 2: Input: [
//   [2147483647,-1,0,2147483647],
//   [2147483647,2147483647,2147483647,-1],
//   [2147483647,-1,2147483647,-1],
//   [0,-1,2147483647,2147483647]
// ]

// Output: [
//   [3,-1,0,1],
//   [2,2,1,-1],
//   [1,-1,2,-1],
//   [0,-1,3,4]
// ]

func isOutOfBounds(r, c int, grid [][]int) bool {
	numRows, numCols := len(grid), len(grid[0])
	if r < 0 || c < 0 {
		return true
	}

	if r >= numRows || c >= numCols {
		return true
	}

	return false
}

func IslandsAndTreasure(grid [][]int) {
	numRows, numCols := len(grid), len(grid[0])
	land := 2147483647
	queue := [][2]int{}
	visited := make([][]bool, numRows)
	var directions = [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for rowIndex := range numRows {
		visited[rowIndex] = make([]bool, numCols)

		for colIndex := range numCols {
			if grid[rowIndex][colIndex] == 0 {
				queue = append(queue, [2]int{rowIndex, colIndex})
			}
		}
	}

	counter := 0

	for len(queue) > 0 {
		counter++
		for range queue {
			first := queue[0]
			queue = queue[1:]
			visited[first[0]][first[1]] = true

			for _, direction := range directions {
				r, c := first[0]+direction[0], first[1]+direction[1]
				if !isOutOfBounds(r, c, grid) && grid[r][c] == land {
					grid[r][c] = counter
					queue = append(queue, [2]int{r, c})
				}
			}
		}
	}
}
