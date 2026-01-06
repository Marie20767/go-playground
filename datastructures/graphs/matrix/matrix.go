package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct {
	data [][]int
}

func New(data [][]int) *Matrix {
	return &Matrix{
		data: data,
	}
}

type Node struct {
	Row    int
	Column int
}

// write a String method for the Matrix that prints it out in a nice way
func (m *Matrix) String() string {
	var sb strings.Builder

	for _, row := range m.data {
		sb.WriteString(fmt.Sprintf("%v\n", row))
	}

	return sb.String()
}

var directions = []Node{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

var directionsWithDiagonals = []Node{
	{-1, 0},  // up
	{0, 1},   // right
	{1, 0},   // down
	{0, -1},  // left
	{-1, 1},  // up-right
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, -1}, // up-left
}

// return 'visitable neighbours' of the given node
// e.g. if you are given Node{Row: 1, Col: 1}: return the nodes above, to the right, to the left, below
// BUT don't return it if it is not 'visitable', e.g. it is 'blocked' or out of bounds
func (m *Matrix) ExploreVisitableNeighbours(node Node) []Node {
	visitable := []Node{}

	for _, direction := range directions {
		nodeToCheck := Node{
			Row:    node.Row + direction.Row,
			Column: node.Column + direction.Column,
		}
		if m.isInBounds(nodeToCheck) && !m.isBlocked(nodeToCheck) {
			visitable = append(visitable, nodeToCheck)
		}
	}

	return visitable
}

func (m *Matrix) ExploreVisitableNeighboursWithDiagonals(node Node) []Node {
	visitable := []Node{}

	for _, direction := range directionsWithDiagonals {
		neighbour := Node{
			Row:    node.Row + direction.Row,
			Column: node.Column + direction.Column,
		}
		if m.isInBounds(neighbour) && !m.isBlocked(neighbour) {
			visitable = append(visitable, neighbour)
		}
	}

	return visitable
}

// returns true if the node is the bottom right node of the matrix
func (m *Matrix) isFinalNode(node Node) bool {
	isLastRow := node.Row == len(m.data)-1
	isLastColumn := node.Column == len(m.data[0])-1

	return isLastRow && isLastColumn
}

// returns false if the node is not inside the matrix bounds
func (m *Matrix) isInBounds(node Node) bool {
	if node.Row < 0 || node.Column < 0 {
		return false
	}

	numRows, numColumns := len(m.data), len(m.data[0])

	if node.Row >= numRows || node.Column >= numColumns {
		return false
	}

	return true
}

// returns true if the node is a blocking node
func (m *Matrix) isBlocked(node Node) bool {
	return m.data[node.Row][node.Column] == 1
}

// starting from the top left node, traverse through the matrix
// return true if you can reach the end, false otherwise
func (m *Matrix) CanReachEndBFS() bool {
	if m.data[0][0] == 1 {
		return false
	}

	queue := []Node{}
	queue = append(queue, Node{Row: 0, Column: 0})
	m.data[0][0] = 1

	for len(queue) > 0 {
		for range len(queue) {
			first := queue[0]
			queue = queue[1:]

			if m.isFinalNode(first) {
				return true
			}

			for _, direction := range directions {
				neighbour := Node{
					Row:    first.Row + direction.Row,
					Column: first.Column + direction.Column,
				}

				if m.isInBounds(neighbour) && m.data[neighbour.Row][neighbour.Column] != 1 && !m.isBlocked(neighbour) {
					queue = append(queue, neighbour)
					m.data[neighbour.Row][neighbour.Column] = 1
				}
			}
		}
	}

	return false
}

// Find the shortest path from top left to bottom right
func (m *Matrix) ShortestPathBFS() int {
	if m.data[0][0] == 1 {
		return 0
	}

	queue := []Node{}
	queue = append(queue, Node{Row: 0, Column: 0})
	counter := 0
	m.data[0][0] = 1

	for len(queue) > 0 {
		for range queue {
			first := queue[0]
			queue = queue[1:]

			if m.isFinalNode(first) {
				return counter
			}

			for _, direction := range directions {
				neighbour := Node{
					Row:    first.Row + direction.Row,
					Column: first.Column + direction.Column,
				}

				if m.isInBounds(neighbour) && m.data[neighbour.Row][neighbour.Column] != 1 && !m.isBlocked(neighbour) {
					m.data[neighbour.Row][neighbour.Column] = 1
					queue = append(queue, neighbour)
				}
			}
		}

		counter++
	}

	return 0
}

// starting from the top left node, traverse through the matrix
// return true if you can reach the end, false otherwise
func (m *Matrix) CanReachEndDFS() bool {
	if m.data[0][0] == 1 {
		return false
	}

	return m.dfs(0, 0)
}

func (m *Matrix) dfs(row int, col int) bool {
	if row == len(m.data)-1 && col == len(m.data[0])-1 {
		return true
	}

	m.data[row][col] = 1
	for _, direction := range directions {
		neighbour := Node{Row: row + direction.Row, Column: col + direction.Column}

		if m.isInBounds(neighbour) && m.data[neighbour.Row][neighbour.Column] != 1 && !m.isBlocked(neighbour) {
			if m.dfs(neighbour.Row, neighbour.Column) {
				return true
			}
		}
	}

	return false
}
