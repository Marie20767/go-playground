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
	numRows := len(m.data) - 1
	if node.Row > numRows || node.Row < 0 {
		return false
	}

	numColumns := len(m.data[node.Row]) - 1

	return node.Column <= numColumns && node.Column > 0
}

// returns true if the node is a blocking node
func (m *Matrix) isBlocked(node Node) bool {
	return m.data[node.Row][node.Column] == 1
}

// ***************************************
// BFS algorithms

// starting from the top left node, traverse through the matrix
// return true if you can reach the end, false otherwise
func (m *Matrix) CanReachEnd() bool {
	queue := []Node{}
	queue = append(queue, Node{Row: 0, Column: 0})
	visited := map[Node]struct{}{
		queue[0]: {},
	}

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

				_, prevVisited := visited[neighbour]
				if !prevVisited && m.isInBounds(neighbour) && !m.isBlocked(neighbour) {
					queue = append(queue, neighbour)
					visited[neighbour] = struct{}{}
				}
			}
		}
	}

	return false
}

func (m *Matrix) CanReachEnd2() bool {
	queue := []Node{}
	queue = append(queue, Node{Row: 0, Column: 0})

	numRows, numCols := len(m.data), len(m.data[0])
	visited := make([][]bool, numRows)
	for rowIndex := range numRows {
		visited[rowIndex] = make([]bool, numCols)
	}

	visited[0][0] = true

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

				if m.isInBounds(neighbour) && !visited[neighbour.Row][neighbour.Column] && !m.isBlocked(neighbour) {
					queue = append(queue, neighbour)
					visited[neighbour.Row][neighbour.Column] = true
				}
			}
		}
	}

	return false
}

// Find the shortest path from top left to bottom right
func (m *Matrix) ShortestPath() int {
	queue := []Node{}
	queue = append(queue, Node{Row: 0, Column: 0})

	numRows, numCols := len(m.data), len(m.data[0])
	visited := make([][]bool, numRows)
	for rowIndex := range visited {
		visited[rowIndex] = make([]bool, numCols)
	}

	visited[0][0] = true
	counter := 0

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

				if m.isInBounds(neighbour) && !visited[neighbour.Row][neighbour.Column] && !m.isBlocked(neighbour) {
					visited[neighbour.Row][neighbour.Column] = true
					queue = append(queue, neighbour)
				}
			}
		}

		counter++
	}

	return 0
}

// ***************************************
