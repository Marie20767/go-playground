package waterflow

// You are given a rectangular island heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

// The islands borders the Pacific Ocean from the top and left sides, and borders the Atlantic Ocean from the bottom and right sides.

// Water can flow in four directions (up, down, left, or right) from a cell to a neighboring cell with height equal or lower. Water can also flow into the ocean from cells adjacent to the ocean.

// Find all cells where water can flow from that cell to both the Pacific and Atlantic oceans. Return it as a 2D list where each element is a list [r, c] representing the row and column of the cell. You may return the answer in any order.

// Examples:

// Example 1:
// Input: heights = [
//   [4,2,7,3,4],
//   [7,4,6,4,7],
//   [6,3,5,3,6]
// ]
// Output: [[0,2],[0,4],[1,0],[1,1],[1,2],[1,3],[1,4],[2,0]]

// Example 2:
// Input: heights = [[1],[1]]
// Output: [[0,0],[0,1]]

type Pair struct {
	r int
	c int
}

var directions = []Pair{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // top
	{-1, 0}, // bottom
}

func isOutOfBounds(r, c int, heights [][]int) bool {
	numRows, numCols := len(heights), len(heights[0])

	if r < 0 || c < 0 {
		return true
	}

	if r >= numRows || c >= numCols {
		return true
	}

	return false
}

func PacificAtlantic(heights [][]int) [][]int {
	h := [][]int{}
	numRows, numCols := len(heights), len(heights[0])
	var dfs func(r, c int, boundary Pair) bool
	visited := getNewVisitedMap(numRows, numCols)

	dfs = func(r, c int, boundary Pair) bool {
		visited[r][c] = true
		if r == boundary.r || c == boundary.c {
			return true
		}

		for _, direction := range directions {
			neighbour := Pair{r + direction.r, c + direction.c}

			if !isOutOfBounds(neighbour.r, neighbour.c, heights) && !visited[neighbour.r][neighbour.c] && heights[r][c] >= heights[neighbour.r][neighbour.c] {
				if dfs(neighbour.r, neighbour.c, boundary) {
					return true
				}
			}
		}

		return false
	}

	for rowIndex := range numRows {
		for colIndex := range numCols {
			canReachPacific := dfs(rowIndex, colIndex, Pair{0, 0})
			visited = getNewVisitedMap(numRows, numCols)
			canReachAtlantic := dfs(rowIndex, colIndex, Pair{numRows - 1, numCols - 1})
			visited = getNewVisitedMap(numRows, numCols)

			if canReachPacific && canReachAtlantic {
				h = append(h, []int{rowIndex, colIndex})
			}
		}
	}

	return h
}

func getNewVisitedMap(numRows, numCols int) [][]bool {
	visited := make([][]bool, numRows)
	for rowIndex := range visited {
		visited[rowIndex] = make([]bool, numCols)
	}

	return visited
}

// better time complexity, not revisiting heights
func PacificAtlanticV2(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)

	var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)
	dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
		coord := [2]int{r, c}
		if visit[coord] ||
			r < 0 || c < 0 ||
			r == rows || c == cols ||
			heights[r][c] < prevHeight {
			return
		}

		visit[coord] = true

		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	for c := range cols {
		dfs(0, c, pac, heights[0][c])
		dfs(rows-1, c, atl, heights[rows-1][c])
	}

	for r := range rows {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols-1, atl, heights[r][cols-1])
	}

	result := make([][]int, 0)
	for r := range rows {
		for c := range cols {
			coord := [2]int{r, c}
			if pac[coord] && atl[coord] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
