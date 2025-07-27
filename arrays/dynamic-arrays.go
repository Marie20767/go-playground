package main

import (
	"fmt"
)

func main() {
	// newArray := newDynamicArray(1)
	// fmt.Println("new array", newArray)

	// newArray.pushback(2)
	// newArray.pushback(3)

	// const index = 1
	// fmt.Printf("element at index %d: %v\n", index, newArray.get(index))

	// fmt.Println("capacity:", newArray.getCapacity())

	nums := []int{1, 2, 1}
	ans := getConcatenation(nums)
	fmt.Println("ans", ans)

}

// challenge 1 - construct dynamic array + methods

type dynamicArray struct {
	capacity int
	size     int
	data     []int
}

func newDynamicArray(capacity int) *dynamicArray {
	return &dynamicArray{
		capacity: capacity,
		size:     0,
		data:     make([]int, 0, capacity),
	}
}

func (array *dynamicArray) get(index int) int {
	return array.data[index]
}

func (array *dynamicArray) getSize() int {
	return array.size
}

func (array *dynamicArray) pushback(element int) {
	if array.size == array.capacity {
		array.resize()
	}

	array.data = append(array.data, element)
	array.size++
}

func (array *dynamicArray) popback() int {
	lastElement := array.data[array.size-1]

	array.data = append(array.data[:array.size-1])
	array.size--

	return lastElement
}

func (array *dynamicArray) resize() {
	array.capacity *= 2

	newData := make([]int, array.size, array.capacity)
	copy(newData, array.data)
	array.data = newData
}

func (array *dynamicArray) getCapacity() int {
	return array.capacity
}

func (array *dynamicArray) set(i int, n int) {
	array.data[i] = n
}

// challenge 2 - concatenation

func getConcatenation(nums []int) []int {
	ans := nums // pointing to the same underlying array
	ans = append(ans, nums...)

	return ans
}