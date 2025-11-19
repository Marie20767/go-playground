package sort

// pick pivot => last element
// set pivotBoundary to 0 (a counter that keeps track of index where everything left of it is less than pivot)
// iterate through array until we get to pivot:
// - if current element less than pivot:
//   - swap it with element at pivot boundary
//   - increment pivotBoundary
// swap pivot with pivotBoundary

// do quicksort on everything left of pivot boundary
// do quicksort on everything right of pivot boundary
// combine them together

// base case - if array size <= 1 then return

func QuickSortNotInPlace(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	pivot := n[len(n)-1]
	pivotBoundary := 0

	for i := 0; i < len(n)-1; i++ {
		if n[i] <= pivot {
			n[pivotBoundary], n[i] = n[i], n[pivotBoundary]
			pivotBoundary++
		}
	}

	n[pivotBoundary], n[len(n)-1] = pivot, n[pivotBoundary]

	left, right := n[:pivotBoundary], n[pivotBoundary+1:]
	sorted := QuickSortNotInPlace(left)
	sorted = append(sorted, pivot)

	return append(sorted, QuickSortNotInPlace(right)...)
}

func QuickSort(n []int) {
	quickSortHelper(n, 0, len(n)-1)
}

func quickSortHelper(n []int, start, end int) {
	if start >= end {
		return
	}

	pivot := n[end]
	pivotBoundary := start

	for i := start; i < end; i++ {
		if n[i] <= pivot {
			n[pivotBoundary], n[i] = n[i], n[pivotBoundary]
			pivotBoundary++
		}
	}

	n[pivotBoundary], n[end] = pivot, n[pivotBoundary]

	quickSortHelper(n, start, pivotBoundary-1)
	quickSortHelper(n, pivotBoundary+1, end)	
}
