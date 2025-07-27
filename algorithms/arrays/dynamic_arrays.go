package dynamic_arrays

type DynamicArray struct {
	capacity int
	size     int
	data     []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity: capacity,
		size:     0,
		data:     make([]int, 0, capacity),
	}
}

func (array *DynamicArray) Get(index int) int {
	return array.data[index]
}

func (array *DynamicArray) GetSize() int {
	return array.size
}

func (array *DynamicArray) Pushback(element int) {
	if array.size == array.capacity {
		array.Resize()
	}

	array.data = append(array.data, element)
	array.size++
}

func (array *DynamicArray) Popback() int {
	lastElement := array.data[array.size-1]

	array.data = array.data[:array.size-1]
	array.size--

	return lastElement
}

func (array *DynamicArray) Resize() {
	array.capacity *= 2

	newData := make([]int, array.size, array.capacity)
	copy(newData, array.data)
	array.data = newData
}

func (array *DynamicArray) GetCapacity() int {
	return array.capacity
}

func (array *DynamicArray) Set(i int, n int) {
	array.data[i] = n
}